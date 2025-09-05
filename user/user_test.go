package user_test

import (
	"github.com/opdss/go-easemob-im-server-sdk/easemob"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

var em *easemob.Easemob

func newEasemob() *easemob.Easemob {
	if em != nil {
		return em
	}
	var err error
	em, err = easemob.NewEasemob(request.Config{
		Endpoints:    []string{"http://baiaaaadu.com/cff", "https://a71.easemob.com"},
		OrgName:      "",
		AppName:      "",
		ClientId:     "",
		ClientSecret: "",
	})
	if err != nil {
		panic(err)
	}
	return em
}
