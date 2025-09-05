package chatroom_test

import (
	"context"
	"fmt"
	"testing"
)

func TestEaseMobChatRoom_GetMembers(t *testing.T) {
	data, err := newEaseMob().Chatroom().GetMembers(context.Background(), "289695344951298", 1, 220)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobChatRoom_RemoveMember(t *testing.T) {
	data, err := newEaseMob().Chatroom().RemoveMember(context.Background(), "289695344951298", "8888881")
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
