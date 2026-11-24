package message

import (
	"context"
)

type ChatgroupMsgReq struct {
	From      *string  `json:"from,omitempty"` //消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	To        []string `json:"to"`
	Type      Type     `json:"type"`                 // 消息类型
	RouteType *string  `json:"route_type,omitempty"` // 路由类型，默认空
}

type SendToChatgroupMsgReq[T MsgBody] struct {
	ChatgroupMsgReq
	Body T `json:"body"`
}

func (m *message) SendTxtToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[TxtMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendImgToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[ImgMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendAudioToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[AudioMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendVideoToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[VideoMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendFileToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[FileMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendLocToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[LocMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendCmdToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[CmdMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendCustomToChatgroup(ctx context.Context, req *SendToChatgroupMsgReq[CustomMsg]) (*SendMsgResp, error) {
	return m.SendToChatgroup(ctx, req)
}

func (m *message) SendToChatgroup(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/chatgroups", req, &resp)
	return &resp, err
}
