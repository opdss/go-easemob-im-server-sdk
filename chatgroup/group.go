package chatgroup

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type Info struct {
	Id                string        `json:"id"`
	Name              string        `json:"name"`
	Avatar            string        `json:"avatar"`
	Description       string        `json:"description"`
	MembersOnly       bool          `json:"membersonly"`
	AllowInvites      bool          `json:"allowinvites"`
	MaxUsers          int           `json:"maxusers"`
	Affiliations      []Affiliation `json:"affiliations"`
	Owner             string        `json:"owner"`
	Created           int64         `json:"created"`
	AffiliationsCount int           `json:"affiliations_count"`
	Disabled          bool          `json:"disabled"`
	Mute              bool          `json:"mute"`
	Public            bool          `json:"public"`
	Custom            string        `json:"custom"`
}
type InfoResp struct {
	request.CommonResp
	Data []Info `json:"data"`
}

type Affiliation struct {
	Member string `json:"member,omitempty"`
	Owner  string `json:"owner,omitempty"`
}

func (c *chatgroup) GetInfo(ctx context.Context, chatgroupId string) (*InfoResp, error) {
	resp := InfoResp{}
	err := c.client.Get(ctx, "/chatgroups/"+chatgroupId, &resp)
	return &resp, err
}

type CreateReq struct {
	GroupId           *string  `json:"groupid,omitempty"`
	GroupName         string   `json:"groupname,omitempty"`
	Avatar            string   `json:"avatar,omitempty"`
	Description       string   `json:"description,omitempty"`
	Public            bool     `json:"public"`
	MaxUsers          *int     `json:"maxusers,omitempty"`
	AllowInvites      *bool    `json:"allowinvites,omitempty"`
	MembersOnly       *bool    `json:"membersonly,omitempty"`
	InviteNeedConfirm *bool    `json:"invite_need_confirm,omitempty"`
	Owner             string   `json:"owner"`
	Members           []string `json:"members,omitempty"`
	Custom            *string  `json:"custom,omitempty"`
}

type CreateResp struct {
	request.CommonResp
	Data struct {
		GroupId string `json:"groupid"`
	} `json:"data"`
}

// Create 创建聊天群
// @document https://doc.easemob.com/document/server-side/group_create.html
func (c *chatgroup) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	resp := CreateResp{}
	err := c.client.Post(ctx, "/chatgroups", req, &resp)
	return &resp, err
}

type UpdateReq struct {
	GroupName         *string `json:"groupname,omitempty"`
	Avatar            *string `json:"avatar,omitempty"`
	Description       *string `json:"description,omitempty"`
	Public            *bool   `json:"public,omitempty"`
	MaxUsers          *int    `json:"maxusers,omitempty"`
	AllowInvites      *bool   `json:"allowinvites,omitempty"`
	MembersOnly       *bool   `json:"membersonly,omitempty"`
	InviteNeedConfirm *bool   `json:"invite_need_confirm,omitempty"`
	Custom            *string `json:"custom,omitempty"`
}

type UpdateResp struct {
	request.CommonResp
	Data struct {
		GroupName         bool `json:"groupname"`
		Avatar            bool `json:"avatar"`
		Description       bool `json:"description,omitempty"`
		Public            bool `json:"public"`
		MaxUsers          bool `json:"maxusers"`
		AllowInvites      bool `json:"allowinvites"`
		MembersOnly       bool `json:"membersonly"`
		InviteNeedConfirm bool `json:"invite_need_confirm"`
		Custom            bool `json:"custom"`
	} `json:"data"`
}

// Update 更新聊天群
// @document https://doc.easemob.com/document/server-side/group_modify.html
func (c *chatgroup) Update(ctx context.Context, chatgroupId string, req *UpdateReq) (*UpdateResp, error) {
	resp := UpdateResp{}
	err := c.client.Put(ctx, "/chatgroups/"+chatgroupId, req, &resp)
	return &resp, err
}

type DeleteResp struct {
	request.CommonResp
	Data struct {
		Success bool   `json:"success"` //是否解散成功
		GroupId string `json:"groupid"`
	} `json:"data"`
}

func (c *chatgroup) Delete(ctx context.Context, chatgroupId string) (*DeleteResp, error) {
	resp := DeleteResp{}
	err := c.client.Delete(ctx, "/chatgroups/"+chatgroupId, nil, &resp)
	return &resp, err
}
