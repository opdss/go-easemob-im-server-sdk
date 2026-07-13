package user

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type BanResp struct {
	request.CommonResp
	Entities []Entity `json:"entities"`
}

// Ban 封禁用户 https://doc.easemob.com/document/server-side/account_ban.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key，详见
func (u *user) Ban(ctx context.Context, username string) (*BanResp, error) {
	resp := &BanResp{}
	err := u.client.Post(ctx, "/users/"+username+"/deactivate", nil, &resp)
	return resp, err
}

// Unban 解禁用户 https://doc.easemob.com/document/server-side/account_unban.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key，详见
func (u *user) Unban(ctx context.Context, username string) (*BanResp, error) {
	resp := &BanResp{}
	err := u.client.Post(ctx, "/users/"+username+"/activate", nil, &resp)
	return resp, err
}
