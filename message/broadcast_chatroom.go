package message

import (
	"context"
)

func (m *message) BroadcastTxtChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastImgChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastAudioChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastVideoChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastFileChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastLocChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCmdChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCustomChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastChatroom(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/chatrooms/broadcast", req, &resp)
	return &resp, err
}
