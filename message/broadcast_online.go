package message

import (
	"context"
)

func (m *message) BroadcastTxtOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastImgOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastAudioOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastVideoOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastFileOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastLocOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCmdOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastCustomOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error) {
	if req.TargetType != nil {
		req.TargetType = nil
	}
	if req.ChatroomMsgLevel != nil {
		req.ChatroomMsgLevel = nil
	}
	return m.Broadcast(ctx, req)
}

func (m *message) BroadcastOnline(ctx context.Context, req any) (*SendMsgResp, error) {
	resp := SendMsgResp{}
	err := m.client.Post(ctx, "/messages/users/broadcast", req, &resp)
	return &resp, err
}
