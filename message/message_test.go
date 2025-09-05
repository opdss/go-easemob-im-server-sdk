package message_test

import (
	"github.com/opdss/go-easemob-im-server-sdk/easemob"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

var em *easemob.Easemob

func newEasemob() *easemob.Easemob {
	if em != nil {
		return em
	}
	em, _ = easemob.NewEasemob(request.Config{
		Endpoints:    []string{"https://a71.easemob.com"},
		OrgName:      "",
		AppName:      "",
		ClientId:     "",
		ClientSecret: "",
	})
	return em
}
