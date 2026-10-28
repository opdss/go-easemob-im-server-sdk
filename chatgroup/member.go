package chatgroup

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
	"strconv"
	"strings"
)

type BatchAddMemberResp struct {
	request.CommonResp
	Data struct {
		NewMembers []string `json:"newmembers"`
		GroupId    string   `json:"groupid"`
		Action     string   `json:"action"`
	} `json:"data"`
}

func (c *chatgroup) BatchAddMembers(ctx context.Context, chatgroupId string, usernames []string, needNotify bool) (*BatchAddMemberResp, error) {
	resp := BatchAddMemberResp{}
	err := c.client.Post(ctx,
		"/chatgroups/"+chatgroupId+"/users?need_notify="+strconv.FormatBool(needNotify),
		map[string]any{"usernames": usernames}, &resp)
	return &resp, err
}

type AddMemberResp struct {
	request.CommonResp
	Data struct {
		Result  bool   `json:"result"`
		GroupId string `json:"groupid"`
		Action  string `json:"action"`
		User    string `json:"user"`
	} `json:"data"`
}

func (c *chatgroup) AddMember(ctx context.Context, chatgroupId string, username string, needNotify bool) (*AddMemberResp, error) {
	resp := AddMemberResp{}
	err := c.client.Post(ctx,
		"/chatgroups/"+chatgroupId+"/users/"+username+"?need_notify="+strconv.FormatBool(needNotify), nil, &resp)
	return &resp, err
}

type MemberResp struct {
	Action      string `json:"action"`
	Application string `json:"application"`
	Params      struct {
		JoinedTime []string `json:"joined_time"`
		PageSize   []string `json:"pagesize"`
		PageNum    []string `json:"pagenum"`
	} `json:"params"`
	Uri             string              `json:"uri"`
	Entities        []interface{}       `json:"entities"`
	Data            []map[string]string `json:"data"`
	Timestamp       int64               `json:"timestamp"`
	Duration        int                 `json:"duration"`
	Organization    string              `json:"organization"`
	ApplicationName string              `json:"applicationName"`
	Count           int                 `json:"count"`
}

func (c *chatgroup) GetMembers(ctx context.Context, chatgroupId string, page, pageSize int, joinedTime bool) (*MemberResp, error) {
	resp := MemberResp{}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 1000
	}
	err := c.client.Get(ctx, "/chatgroups/"+chatgroupId+"/users?pagenum="+strconv.Itoa(page)+"&pagesize="+strconv.Itoa(pageSize)+"&joined_time="+strconv.FormatBool(joinedTime), &resp)
	return &resp, err
}

type RemoveMemberRespItem struct {
	Result  bool   `json:"result"`
	Action  string `json:"action"`
	User    string `json:"user"`
	GroupId string `json:"groupid"`
}

type RemoveMemberResp struct {
	request.CommonResp
	Data RemoveMemberRespItem `json:"data"`
}

func (c *chatgroup) RemoveMember(ctx context.Context, chatgroupId string, username string, needNotify bool) (*RemoveMemberResp, error) {
	resp := RemoveMemberResp{}
	err := c.client.Delete(ctx, "/chatgroups/"+chatgroupId+"/users/"+username+"?need_notify="+strconv.FormatBool(needNotify), nil, &resp)
	return &resp, err
}

type BatchRemoveMemberResp struct {
	request.CommonResp
	Data []RemoveMemberRespItem `json:"data"`
}

func (c *chatgroup) BatchRemoveMember(ctx context.Context, chatgroupId string, usernames []string, needNotify bool) (*BatchRemoveMemberResp, error) {
	resp := BatchRemoveMemberResp{}
	userStr := strings.Join(usernames, ",")
	err := c.client.Delete(ctx, "/chatgroups/"+chatgroupId+"/users/"+userStr+"?need_notify="+strconv.FormatBool(needNotify), nil, &resp)
	return &resp, err
}
