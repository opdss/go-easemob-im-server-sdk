package message_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/message"
	"testing"
)

func TestMessage_Broadcast(t *testing.T) {

	data, err := newEasemob().Message().BroadcastTxt(context.Background(), &message.BroadcastMsgReq[message.BroadcastTxtMsg]{
		TargetType: "users",
		//Appkey:     newEasemob().Client.Config().AppKey,
		//From: utils.PointerAny("888888"),
		Msg: message.BroadcastTxtMsg{
			Type: message.Txt,
			TxtMsg: message.TxtMsg{
				Msg: "test111",
			},
		},
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

}
