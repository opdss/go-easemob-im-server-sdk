package user

import (
	"context"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

type DisconnectResp struct {
	request.CommonResp
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
}

// Disconnect 强制将用户下线 https://doc.easemob.com/document/server-side/account_offline_forced.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key
func (u *user) Disconnect(ctx context.Context, username string) (*DisconnectResp, error) {
	resp := &DisconnectResp{}
	err := u.client.Get(ctx, "/users/"+username+"/disconnect", &resp)
	return resp, err
}

// DisconnectSingleDevice 强制将用户单设备下线 https://doc.easemob.com/document/server-side/account_offline_device_single.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key
func (u *user) DisconnectSingleDevice(ctx context.Context, username string, resourceId string) (*DisconnectResp, error) {
	resp := &DisconnectResp{}
	err := u.client.Delete(ctx, "/users/"+username+"/disconnect/"+resourceId, nil, &resp)
	return resp, err
}

type OnlineStatusResp struct {
	request.CommonResp
	Data map[string]string `json:"data"`
}

// GetOnlineStatus 获取用户在线状态 https://doc.easemob.com/document/server-side/account_presence_obtain_single.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key
func (u *user) GetOnlineStatus(ctx context.Context, username string) (*OnlineStatusResp, error) {
	resp := &OnlineStatusResp{}
	err := u.client.Get(ctx, "/users/"+username+"/status", &resp)
	return resp, err
}

type BatchOnlineStatusResp struct {
	request.CommonResp
	Data []map[string]string `json:"data"`
}

// BatchGetOnlineStatus 批量获取用户在线状态 https://doc.easemob.com/document/server-side/account_presence_obtain_batch.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key
func (u *user) BatchGetOnlineStatus(ctx context.Context, usernames []string) (*BatchOnlineStatusResp, error) {
	resp := &BatchOnlineStatusResp{}
	err := u.client.Post(ctx, "/users/batch/status", map[string][]string{"usernames": usernames}, &resp)
	return resp, err
}

type OnlineDevicesResp struct {
	request.CommonResp
	Data []OnlineDevice `json:"data"`
}

type OnlineDevice struct {
	Res        string `json:"res"`
	DeviceUuid string `json:"device_uuid"`
	DeviceName string `json:"device_name"`
}

// GetOnlineDevices 获取用户在线状态 https://doc.easemob.com/document/server-side/account_online_device_obtain.html
// 该 API、用户账户管理的其他接口、以及离线推送的相关接口的总调用频率上限为 100 次/秒/App Key
func (u *user) GetOnlineDevices(ctx context.Context, username string) (*OnlineDevicesResp, error) {
	resp := &OnlineDevicesResp{}
	err := u.client.Get(ctx, "/users/"+username+"/resources", &resp)
	return resp, err
}
