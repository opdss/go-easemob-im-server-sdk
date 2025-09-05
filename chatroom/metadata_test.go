package chatroom_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/chatroom"
	"testing"
)

func TestEaseMobChatRoom_GetAnnouncement(t *testing.T) {
	data, err := newEaseMob().Chatroom().GetAnnouncement(context.Background(), "289695344951298")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_SetAnnouncement(t *testing.T) {
	data, err := newEaseMob().Chatroom().SetAnnouncement(context.Background(), "289695344951298", "888888的地盘")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_SetMetadata(t *testing.T) {
	data, err := newEaseMob().Chatroom().SetMetadata(context.Background(), "289695344951298", "888888", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_GetMetadata(t *testing.T) {
	data, err := newEaseMob().Chatroom().GetMetadata(context.Background(), "289695344951298", []string{"key", "key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_DelMetadata(t *testing.T) {
	data, err := newEaseMob().Chatroom().DelMetadata(context.Background(), "289695344951298", "888888", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_ForceSetMetadata(t *testing.T) {
	data, err := newEaseMob().Chatroom().ForceSetMetadata(context.Background(), "289695344951298", "8888881", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_ForceDelMetadata(t *testing.T) {
	data, err := newEaseMob().Chatroom().DelMetadata(context.Background(), "289695344951298", "8888881", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
