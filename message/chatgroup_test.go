package message_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/message"
	"testing"
)

func TestMessage_SendTxtToChatgroup(t *testing.T) {

	data, err := newEasemob().Message().SendTxtToChatgroup(context.Background(), &message.SendToChatgroupMsgReq[message.TxtMsg]{
		ChatgroupMsgReq: message.ChatgroupMsgReq{
			To:   []string{"88884488"},
			Type: message.Txt,
		},
		Body: message.TxtMsg{
			Msg: "xxxxx",
		},
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

}
