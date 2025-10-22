package request

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"golang.org/x/sync/singleflight"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Endpoints              []string `help:"环信接口，可以配置多个，会自动切换, 这里注意后缀不可以带有/" default:""`
	OrgName                string   `help:"环信企业名" default:""`
	AppName                string   `help:"环信应用名" default:""`
	ClientId               string   `help:"环信应用clientId" default:""`
	ClientSecret           string   `help:"环信应用clientSecret" default:""`
	TokenTTL               int64    `help:"环信token有效期，单位秒" default:"86400"`
	ChangeEndpointInterval int64    `help:"环信切换endpoint的时间间隔，单位秒" default:"10"`
	Debug                  bool     `help:"是否开启debug模式，开启后会打印请求和响应" default:"false" devDefault:"true"`
}

type CommonResp struct {
	Path            string `json:"path"`
	Uri             string `json:"uri"`
	Timestamp       int64  `json:"timestamp"` //毫秒
	Organization    string `json:"organization"`
	Application     string `json:"application"`
	Action          string `json:"action"`
	Duration        int    `json:"duration"` //执行时间，ms
	ApplicationName string `json:"applicationName"`
}
type ErrorResp struct {
	Code             int    `json:"-"`
	Error_           string `json:"error"`
	Exception        string `json:"exception"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration"`
	ErrorDescription string `json:"error_description"`
}

func (e *ErrorResp) Error() string {
	return fmt.Sprintf("%s[%d]: %s", e.Error_, e.Code, e.ErrorDescription)
}

type RequestErrorResp struct {
	Code int    `json:"code"`
	Body string `json:"body"`
	Err  error  `json:"err"`
}

func (e *RequestErrorResp) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("%d: %s", e.Code, e.Body)
}

type TokenResp struct {
	AccessToken string    `json:"access_token"` //有效的 Token 字符串。
	Application string    `json:"token_type"`   //当前 App 的 UUID 值。
	ExpiresIn   int64     `json:"expires_in"`   //Token 有效时间，单位为秒，在有效期内不需要重复获取。
	ExpiresAt   time.Time `json:"-"`            //Token 有效截止时间，在有效期内不需要重复获取。接口不返回，用于计算
}

type Client struct {
	config                 Config
	token                  *TokenResp
	singleFlight           *singleflight.Group
	endpointLock           *sync.RWMutex
	lastChangeEndpointTime int64 // 上次切换 endpoint 的时间
}

func NewClient(conf Config) (*Client, error) {
	if len(conf.Endpoints) == 0 || conf.AppName == "" || conf.OrgName == "" || conf.ClientId == "" || conf.ClientSecret == "" {
		return nil, errors.New("环信配置错误，请检查")
	}
	return &Client{
		config:       conf,
		singleFlight: &singleflight.Group{},
		endpointLock: &sync.RWMutex{},
	}, nil
}

// Config 获取当前配置
func (c *Client) Config() Config {
	return c.config
}

// Get 发送 GET 请求
func (c *Client) Get(ctx context.Context, apiPath string, resp any) error {
	return c.requestWithToken(ctx, http.MethodGet, apiPath, nil, resp)
}

// Post 发送 POST 请求
func (c *Client) Post(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodPost, apiPath, req, resp)
}

// Put 发送 PUT 请求
func (c *Client) Put(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodPut, apiPath, req, resp)
}

// Delete 发送 DELETE 请求
func (c *Client) Delete(ctx context.Context, apiPath string, req any, resp any) error {
	return c.requestWithToken(ctx, http.MethodDelete, apiPath, req, resp)
}

// requestWithToken 发送请求并携带 token
func (c *Client) requestWithToken(ctx context.Context, method string, apiPath string, req any, resp any) error {
	tk, err := c.GetToken(ctx)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if req != nil {
		err = json.NewEncoder(buf).Encode(req)
		if err != nil {
			return err
		}
	}
	if c.config.Debug {
		log.Printf("easemob[client.requestWithToken] request: %s => %s\n", apiPath, buf.String())
	}
	httpReq, err := http.NewRequestWithContext(ctx, method, c.getApiUrl(apiPath), buf)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Authorization", "Bearer "+tk)
	return c.request(httpReq, resp)
}

// request 发送请求
func (c *Client) request(req *http.Request, resp any) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		if utils.IsNetError(err) {
			//切域名
			c.changeEndpoint()
		}
		return err
	}
	body, err := io.ReadAll(httpResp.Body)
	defer func() {
		_ = httpResp.Body.Close()
	}()
	if c.config.Debug {
		log.Println("easemob[client.request] response: ", string(body))
	}
	if err != nil {
		return &RequestErrorResp{Code: httpResp.StatusCode, Err: err, Body: string(body)}
	}
	if httpResp.StatusCode != 200 {
		errResp := &ErrorResp{}
		err = json.Unmarshal(body, errResp)
		if err != nil {
			return &RequestErrorResp{Code: httpResp.StatusCode, Err: err, Body: string(body)}
		}
		errResp.Code = httpResp.StatusCode
		return errResp
	}
	return json.Unmarshal(body, &resp)
}

// getApiUrl 获取 api 地址
func (c *Client) getApiUrl(path string) string {
	c.endpointLock.RLock()
	defer c.endpointLock.RUnlock()
	return fmt.Sprintf("%s/%s/%s%s", c.config.Endpoints[0], c.config.OrgName, c.config.AppName, path)
}

// GetToken 获取授权 token
func (c *Client) GetToken(ctx context.Context) (string, error) {
	_resp, err, _ := c.singleFlight.Do("GetToken", func() (any, error) {
		if c.token != nil && c.token.AccessToken != "" && c.token.ExpiresAt.After(time.Now()) {
			return c.token.AccessToken, nil
		}
		resp, err := c.getToken(ctx)
		if err != nil {
			return "", err
		}
		c.token = resp
		c.token.ExpiresAt = time.Now().Add(time.Second*time.Duration(resp.ExpiresIn) - 30) //预留30秒
		return resp.AccessToken, nil
	})
	if err != nil {
		return "", err
	}
	return _resp.(string), nil
}

// getToken 获取 token
func (c *Client) getToken(ctx context.Context) (*TokenResp, error) {
	data := map[string]any{
		"grant_type":    "client_credentials",
		"client_id":     c.config.ClientId,
		"client_secret": c.config.ClientSecret,
		"ttl":           c.config.TokenTTL,
	}
	if c.config.TokenTTL > 0 {
		data["ttl"] = c.config.TokenTTL
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.getApiUrl("/token"), &buf)
	if err != nil {
		return nil, err
	}
	resp := &TokenResp{}
	err = c.request(req, resp)
	if c.config.Debug {
		log.Printf("easemob[client.getToken] get new token request: %s\n", buf.String())
		if err == nil {
			_s, _ := json.Marshal(resp)
			log.Printf("easemob[client.getToken] get new token reponse: %s\n", string(_s))
		} else {
			log.Println("easemob[client.getToken] get new token error: ", err.Error())
		}
	}
	return resp, err
}

// ChangeEndpoint 自动切换 Api 服务器地址
func (c *Client) changeEndpoint() {
	if len(c.config.Endpoints) < 2 {
		return
	}
	// 检查距离上次更换uri的时间间隔
	c.endpointLock.Lock()
	defer c.endpointLock.Unlock()
	nowUnix := time.Now().Unix()
	if (nowUnix - c.lastChangeEndpointTime) >= c.config.ChangeEndpointInterval {
		c.config.Endpoints = append(c.config.Endpoints[1:], c.config.Endpoints[0])
		c.lastChangeEndpointTime = nowUnix
	}
}
