package chatroom_test

import (
	"github.com/opdss/go-easemob-im-server-sdk/easemob"
	"github.com/opdss/go-easemob-im-server-sdk/request"
)

var em *easemob.EaseMob

func newEaseMob() *easemob.EaseMob {
	if em != nil {
		return em
	}
	em, _ = easemob.NewEaseMob(request.Config{
		Endpoints:    []string{"https://a71.easemob.com"},
		OrgName:      "1104241128170445",
		AppName:      "sout-test",
		ClientId:     "YXA6OIEyvWabTTyhZBY8tlRRUg",
		ClientSecret: "YXA6iyTY12_6S1XehnmOjqF2DZBWCFg",
	})
	return em
}
