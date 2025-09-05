package easemob

import (
	"github.com/opdss/go-easemob-im-server-sdk/chatroom"
	"github.com/opdss/go-easemob-im-server-sdk/message"
	"github.com/opdss/go-easemob-im-server-sdk/request"
	"github.com/opdss/go-easemob-im-server-sdk/user"
)

type EaseMob struct {
	*request.Client
	chatroom chatroom.Chatroom
	user     user.User
	message  message.Message
}

func NewEaseMob(conf request.Config) (*EaseMob, error) {
	c, err := request.NewClient(conf)
	if err != nil {
		return nil, err
	}
	em := &EaseMob{
		Client:   c,
		chatroom: chatroom.NewChatroom(c),
		user:     user.NewUser(c),
		message:  message.NewMessage(c),
	}
	return em, nil
}

func (em *EaseMob) User() user.User {
	return em.user
}

func (em *EaseMob) Chatroom() chatroom.Chatroom {
	return em.chatroom
}

func (em *EaseMob) Message() message.Message {
	return em.message
}
