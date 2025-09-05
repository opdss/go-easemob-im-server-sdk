package chatroom_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/chatroom"
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"testing"
)

func TestEasemobChatRoom_GetInfo(t *testing.T) {
	data, err := newEasemob().Chatroom().GetInfo(context.Background(), "289695344951298")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_Create(t *testing.T) {
	data, err := newEasemob().Chatroom().Create(context.Background(), &chatroom.CreateReq{
		Name:        "chat room 888888",
		Description: "chatroom Description",
		MaxUsers:    utils.PointerAny(111),
		Owner:       "888888",
		Members:     []string{"8888881"},
		Custom:      utils.PointerAny("chatroom Custom"),
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_Update(t *testing.T) {
	data, err := newEasemob().Chatroom().Update(context.Background(), "289695344951298", &chatroom.UpdateReq{
		Name:        "chat room 289695344951298",
		Description: "chatroom Description",
		MaxUsers:    8888,
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
