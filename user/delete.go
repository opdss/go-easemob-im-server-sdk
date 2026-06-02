package user

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type DeleteResp struct {
	request.CommonResp
	Entities []Entity `json:"entities"`
}

// Delete 删除用户 https://doc.easemob.com/document/server-side/account_system.html#%E6%8E%88%E6%9D%83%E6%B3%A8%E5%86%8C%E5%8D%95%E4%B8%AA%E7%94%A8%E6%88%B7
func (u *user) Delete(ctx context.Context, username string) (*DeleteResp, error) {
	resp := &DeleteResp{}
	err := u.client.Delete(ctx, "/users/"+username, nil, &resp)
	return resp, err
}
