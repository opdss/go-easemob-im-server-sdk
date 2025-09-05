package chatroom_test

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
		OrgName:      "1104241128170445",
		AppName:      "sout-test",
		ClientId:     "YXA6OIEyvWabTTyhZBY8tlRRUg",
		ClientSecret: "YXA6iyTY12_6S1XehnmOjqF2DZBWCFg",
	})
	return em
}
