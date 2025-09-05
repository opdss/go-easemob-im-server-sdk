package user_test

import (
	"github.com/opdss/go-easemob-im-server-sdk/easemob"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

var em *easemob.EaseMob

func newEaseMob() *easemob.EaseMob {
	if em != nil {
		return em
	}
	var err error
	em, err = easemob.NewEaseMob(request.Config{
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
