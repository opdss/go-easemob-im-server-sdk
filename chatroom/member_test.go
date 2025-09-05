package chatroom_test

import (
	"context"
	"fmt"
	"testing"
)

func TestEasemobChatRoom_GetMembers(t *testing.T) {
	data, err := newEasemob().Chatroom().GetMembers(context.Background(), "289695344951298", 1, 220)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobChatRoom_RemoveMember(t *testing.T) {
	data, err := newEasemob().Chatroom().RemoveMember(context.Background(), "289695344951298", "8888881")
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
