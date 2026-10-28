package chatgroup_test

import (
	"context"
	"fmt"
	"testing"
)

func TestEasemobChatGroup_GetMembers(t *testing.T) {
	data, err := newEasemob().Chatgroup().GetMembers(context.Background(), "666777888", 1, 220, false)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
func TestEasemobChatGroup_AddMember(t *testing.T) {
	data, err := newEasemob().Chatgroup().AddMember(context.Background(), "666777888", "101968", false)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_BatchAddMembers(t *testing.T) {
	data, err := newEasemob().Chatgroup().BatchAddMembers(context.Background(), "666777888", []string{"105946", "108949"}, false)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_BatchRemoveMembers(t *testing.T) {
	data, err := newEasemob().Chatgroup().BatchRemoveMember(context.Background(), "666777888", []string{"105946", "108949"}, false)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatGroup_RemoveMember(t *testing.T) {
	data, err := newEasemob().Chatgroup().RemoveMember(context.Background(), "666777888", "101968", false)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
