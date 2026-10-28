package chatgroup_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/chatgroup"
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"testing"
)

func TestEasemobChatGroup_GetInfo(t *testing.T) {
	data, err := newEasemob().Chatgroup().GetInfo(context.Background(), "666777888")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_Create(t *testing.T) {
	data, err := newEasemob().Chatgroup().Create(context.Background(), &chatgroup.CreateReq{
		GroupId:     utils.ToPointer("666777888"),
		GroupName:   "chat group 888888",
		Description: "chatroom Description",
		MaxUsers:    utils.ToPointer(111),
		Owner:       "100619",
		Members:     []string{"100619"},
		Custom:      utils.ToPointer("chatroom Custom"),
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_Update(t *testing.T) {
	data, err := newEasemob().Chatgroup().Update(context.Background(), "666777888", &chatgroup.UpdateReq{
		GroupName:   utils.ToPointer("bbbbbb"),
		Description: utils.ToPointer(""),
		MaxUsers:    utils.ToPointer(300),
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_Delete(t *testing.T) {
	data, err := newEasemob().Chatgroup().Delete(context.Background(), "666777888")

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
