package easemob

import (
	"github.com/opdss/go-easemob-im-server-sdk/chatroom"
	"github.com/opdss/go-easemob-im-server-sdk/message"
	"github.com/opdss/go-easemob-im-server-sdk/request"
	"github.com/opdss/go-easemob-im-server-sdk/user"
)

type Easemob struct {
	*request.Client
	chatroom chatroom.Chatroom
	user     user.User
	message  message.Message
}

func NewEasemob(conf request.Config) (*Easemob, error) {
	c, err := request.NewClient(conf)
	if err != nil {
		return nil, err
	}
	em := &Easemob{
		Client:   c,
		chatroom: chatroom.NewChatroom(c),
		user:     user.NewUser(c),
		message:  message.NewMessage(c),
	}
	return em, nil
}

func (em *Easemob) User() user.User {
	return em.user
}

func (em *Easemob) Chatroom() chatroom.Chatroom {
	return em.chatroom
}

func (em *Easemob) Message() message.Message {
	return em.message
}
