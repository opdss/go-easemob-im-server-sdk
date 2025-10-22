package message

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type Message interface {
	// SendTxtToUser 发送文本消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E6%96%87%E6%9C%AC%E6%B6%88%E6%81%AF
	SendTxtToUser(ctx context.Context, req *SendToUserMsgReq[TxtMsg]) (*SendMsgResp, error)
	// SendImgToUser 发送图片消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E5%9B%BE%E7%89%87%E6%B6%88%E6%81%AF
	SendImgToUser(ctx context.Context, req *SendToUserMsgReq[ImgMsg]) (*SendMsgResp, error)
	// SendAudioToUser 发送音频消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E8%AF%AD%E9%9F%B3%E6%B6%88%E6%81%AF
	SendAudioToUser(ctx context.Context, req *SendToUserMsgReq[AudioMsg]) (*SendMsgResp, error)
	// SendVideoToUser 发送视频消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E8%A7%86%E9%A2%91%E6%B6%88%E6%81%AF
	SendVideoToUser(ctx context.Context, req *SendToUserMsgReq[VideoMsg]) (*SendMsgResp, error)
	// SendFileToUser 发送文件消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E6%96%87%E4%BB%B6%E6%B6%88%E6%81%AF
	SendFileToUser(ctx context.Context, req *SendToUserMsgReq[FileMsg]) (*SendMsgResp, error)
	// SendLocToUser 发送位置消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E4%BD%8D%E7%BD%AE%E6%B6%88%E6%81%AF
	SendLocToUser(ctx context.Context, req *SendToUserMsgReq[LocMsg]) (*SendMsgResp, error)
	// SendCmdToUser 发送透传消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E9%80%8F%E4%BC%A0%E6%B6%88%E6%81%AF
	SendCmdToUser(ctx context.Context, req *SendToUserMsgReq[CmdMsg]) (*SendMsgResp, error)
	// SendCustomToUser 发送自定义消息到用户
	// https://doc.easemob.com/document/server-side/message_single.html#%E5%8F%91%E9%80%81%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
	SendCustomToUser(ctx context.Context, req *SendToUserMsgReq[CustomMsg]) (*SendMsgResp, error)
	// SendToUser 发送消息到用户
	SendToUser(ctx context.Context, req any) (*SendMsgResp, error)

	// SendTxtToChatroom 发送文本消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E6%96%87%E6%9C%AC%E6%B6%88%E6%81%AF
	SendTxtToChatroom(ctx context.Context, req *SendToChatroomMsgReq[TxtMsg]) (*SendMsgResp, error)
	// SendImgToChatroom 发送图片消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E5%9B%BE%E7%89%87%E6%B6%88%E6%81%AF
	SendImgToChatroom(ctx context.Context, req *SendToChatroomMsgReq[ImgMsg]) (*SendMsgResp, error)
	// SendAudioToChatroom 发送音频消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E8%AF%AD%E9%9F%B3%E6%B6%88%E6%81%AF
	SendAudioToChatroom(ctx context.Context, req *SendToChatroomMsgReq[AudioMsg]) (*SendMsgResp, error)
	// SendVideoToChatroom 发送视频消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E8%A7%86%E9%A2%91%E6%B6%88%E6%81%AF
	SendVideoToChatroom(ctx context.Context, req *SendToChatroomMsgReq[VideoMsg]) (*SendMsgResp, error)
	// SendFileToChatroom 发送文件消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E6%96%87%E4%BB%B6%E6%B6%88%E6%81%AF
	SendFileToChatroom(ctx context.Context, req *SendToChatroomMsgReq[FileMsg]) (*SendMsgResp, error)
	// SendLocToChatroom 发送位置消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E4%BD%8D%E7%BD%AE%E6%B6%88%E6%81%AF
	SendLocToChatroom(ctx context.Context, req *SendToChatroomMsgReq[LocMsg]) (*SendMsgResp, error)
	// SendCmdToChatroom 发送透传消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E9%80%8F%E4%BC%A0%E6%B6%88%E6%81%AF
	SendCmdToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CmdMsg]) (*SendMsgResp, error)
	// SendCustomToChatroom 发送自定义消息到聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#%E5%8F%91%E9%80%81%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
	SendCustomToChatroom(ctx context.Context, req *SendToChatroomMsgReq[CustomMsg]) (*SendMsgResp, error)
	// SendToChatroom 发送消息到聊天室
	SendToChatroom(ctx context.Context, req any) (*SendMsgResp, error)

	// BroadcastTxt 广播消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastTxt(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error)
	// BroadcastImg 广播图片消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastImg(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error)
	// BroadcastAudio 广播音频消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastAudio(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error)
	// BroadcastVideo 广播视频消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastVideo(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error)
	// BroadcastFile 广播文件消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastFile(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error)
	// BroadcastLoc 广播位置消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastLoc(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error)
	// BroadcastCmd 广播透传消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCmd(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error)
	// BroadcastCustom 广播自定义消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCustom(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error)
	// Broadcast 广播消息发送所有用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E6%89%80%E6%9C%89%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	Broadcast(ctx context.Context, req any) (*SendMsgResp, error)

	// BroadcastTxtOnline 广播文本消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastTxtOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error)
	// BroadcastImgOnline 广播图片消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastImgOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error)
	// BroadcastAudioOnline 广播音频消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastAudioOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error)
	// BroadcastVideoOnline 广播视频消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastVideoOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error)
	// BroadcastFileOnline 广播文件消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastFileOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error)
	// BroadcastLocOnline 广播位置消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastLocOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error)
	// BroadcastCmdOnline 广播透传消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCmdOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error)
	// BroadcastCustomOnline 广播自定义消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%90%91-app-%E5%9C%A8%E7%BA%BF%E7%94%A8%E6%88%B7%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCustomOnline(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error)
	// BroadcastOnline 广播消息发送在线用户，慎用
	BroadcastOnline(ctx context.Context, req any) (*SendMsgResp, error)

	// BroadcastTxtChatroom 广播文本消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastTxtChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastTxtMsg]) (*SendMsgResp, error)
	// BroadcastImgChatroom 广播图片消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastImgChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastImgMsg]) (*SendMsgResp, error)
	// BroadcastAudioChatroom 广播音频消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastAudioChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastAudioMsg]) (*SendMsgResp, error)
	// BroadcastVideoChatroom 广播视频消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastVideoChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastVideoMsg]) (*SendMsgResp, error)
	// BroadcastFileChatroom 广播文件消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastFileChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastFileMsg]) (*SendMsgResp, error)
	// BroadcastLocChatroom 广播位置消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastLocChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastLocMsg]) (*SendMsgResp, error)
	// BroadcastCmdChatroom 广播透传消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCmdChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastCmdMsg]) (*SendMsgResp, error)
	// BroadcastCustomChatroom 广播自定义消息发送在线用户，慎用
	// https://doc.easemob.com/document/server-side/message_broadcast.html#%E5%8F%91%E9%80%81%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%85%A8%E5%B1%80%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
	BroadcastCustomChatroom(ctx context.Context, req *BroadcastMsgReq[BroadcastCustomMsg]) (*SendMsgResp, error)
	// BroadcastChatroom 广播消息发送在线用户，慎用
	BroadcastChatroom(ctx context.Context, req any) (*SendMsgResp, error)
}

type message struct {
	client *request.Client
}

func NewMessage(client *request.Client) Message {
	return &message{
		client: client,
	}
}

type MsgLevel string

const (
	High   MsgLevel = "high"   //高
	Low    MsgLevel = "low"    //低
	Normal MsgLevel = "normal" //普通
)

// Type 消息类型
type Type string

const (
	Txt    Type = "txt"    //文本消息
	Img    Type = "img"    //图片消息
	Audio  Type = "audio"  //  语音消息
	Video  Type = "video"  //视频消息
	File   Type = "file"   //文件消息
	Loc    Type = "loc"    //位置消息
	Cmd    Type = "cmd"    //透传消息
	Custom Type = "custom" //自定义消息
)

// TxtMsg 文本消息
type TxtMsg struct {
	Msg string `json:"msg"`
}

type ImgMsgSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ImgMsg 图片消息
type ImgMsg struct {
	Filename *string     `json:"filename,omitempty"`
	Secret   *string     `json:"secret,omitempty"`
	Url      string      `json:"url"`
	Size     *ImgMsgSize `json:"size,omitempty"`
}

// AudioMsg 语音消息
type AudioMsg struct {
	Url      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
	Length   *int    `json:"length,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

// VideoMsg 视频消息
type VideoMsg struct {
	Filename    *string `json:"filename,omitempty"`
	Thumb       *string `json:"thumb,omitempty"`
	Length      *int    `json:"length,omitempty"`
	Secret      *string `json:"secret,omitempty"`
	FileLength  *int64  `json:"file_length,omitempty"`
	ThumbSecret *string `json:"thumb_secret,omitempty"`
	Url         string  `json:"url"`
}

// FileMsg 文件消息
type FileMsg struct {
	Filename *string `json:"filename,omitempty"`
	Secret   *string `json:"secret,omitempty"`
	Url      string  `json:"url"`
}

// LocMsg 位置消息
type LocMsg struct {
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
	Addr string `json:"addr"`
}

// CmdMsg 透传消息
type CmdMsg struct {
	Action string `json:"action"`
}

// CustomMsg 自定义消息body
type CustomMsg struct {
	CustomEvent *string    `json:"customEvent,omitempty"`
	CustomExts  CustomExts `json:"customExts,omitempty"`
}

type CustomExts map[string]string

type MsgBody interface {
	TxtMsg | ImgMsg | AudioMsg | VideoMsg | FileMsg | LocMsg | CmdMsg | CustomMsg
}

type MsgData map[string]string

type SendMsgResp struct {
	request.CommonResp
	Data MsgData `json:"data"`
}
