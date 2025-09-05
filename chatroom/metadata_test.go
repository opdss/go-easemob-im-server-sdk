package chatroom_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/chatroom"
	"testing"
)

func TestEasemobChatRoom_GetAnnouncement(t *testing.T) {
	data, err := newEasemob().Chatroom().GetAnnouncement(context.Background(), "289695344951298")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_SetAnnouncement(t *testing.T) {
	data, err := newEasemob().Chatroom().SetAnnouncement(context.Background(), "289695344951298", "888888的地盘")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_SetMetadata(t *testing.T) {
	data, err := newEasemob().Chatroom().SetMetadata(context.Background(), "289695344951298", "888888", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_GetMetadata(t *testing.T) {
	data, err := newEasemob().Chatroom().GetMetadata(context.Background(), "289695344951298", []string{"key", "key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_DelMetadata(t *testing.T) {
	data, err := newEasemob().Chatroom().DelMetadata(context.Background(), "289695344951298", "888888", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_ForceSetMetadata(t *testing.T) {
	data, err := newEasemob().Chatroom().ForceSetMetadata(context.Background(), "289695344951298", "8888881", &chatroom.SetMetaDataReq{
		MetaData: map[string]string{
			"key":  "value",
			"key2": "value2",
		},
		AutoDelete: chatroom.NoDelete,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_ForceDelMetadata(t *testing.T) {
	data, err := newEasemob().Chatroom().DelMetadata(context.Background(), "289695344951298", "8888881", []string{"key2", "key3"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
