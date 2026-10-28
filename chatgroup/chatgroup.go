package chatgroup

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type Chatgroup interface {
	// GetInfo 获取聊天室信息
	// https://doc.easemob.com/document/server-side/group_manage.html#%E6%9F%A5%E8%AF%A2%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%AF%A6%E6%83%85
	GetInfo(ctx context.Context, chatgroupId string) (*InfoResp, error)
	// Create 创建聊天群组
	// https://doc.easemob.com/document/server-side/group_create.html
	Create(ctx context.Context, req *CreateReq) (*CreateResp, error)
	// Update 修改聊天群组信息
	// https://doc.easemob.com/document/server-side/group_modify.html
	Update(ctx context.Context, chatgroupId string, req *UpdateReq) (*UpdateResp, error)
	// Delete 删除聊天群组
	// https://doc.easemob.com/document/server-side/group_delete.html
	Delete(ctx context.Context, chatgroupId string) (*DeleteResp, error)

	// BatchAddMembers 批量添加群组成员
	// https://doc.easemob.com/document/server-side/group_member_add_batch.html
	BatchAddMembers(ctx context.Context, chatgroupId string, usernames []string, needNotify bool) (*BatchAddMemberResp, error)
	// AddMember 添加群组成员
	// https://doc.easemob.com/document/server-side/group_member_add_single.html
	AddMember(ctx context.Context, chatgroupId string, username string, needNotify bool) (*AddMemberResp, error)
	// GetMembers 获取群组成员列表
	// https://doc.easemob.com/document/server-side/group_member_list_obtain.html
	GetMembers(ctx context.Context, chatgroupId string, page, pageSize int, joinedTime bool) (*MemberResp, error)
	// RemoveMember 移除群组成员
	// https://doc.easemob.com/document/server-side/group_member_remove_single.html
	RemoveMember(ctx context.Context, chatgroupId string, username string, needNotify bool) (*RemoveMemberResp, error)
	// BatchRemoveMember 批量移除群组成员
	// https://doc.easemob.com/document/server-side/group_member_remove_batch.html
	BatchRemoveMember(ctx context.Context, chatgroupId string, usernames []string, needNotify bool) (*BatchRemoveMemberResp, error)
}

type chatgroup struct {
	client request.Client
}

func NewChatgroup(client request.Client) Chatgroup {
	return &chatgroup{
		client: client,
	}
}
