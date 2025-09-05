package user_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/user"
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"testing"
)

func TestEaseMobUser_Register(t *testing.T) {
	data, err := newEaseMob().User().Register(context.Background(), &user.RegisterReq{Username: "8888881", Password: "123456"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobUser_BatchRegistry(t *testing.T) {
	data, err := newEaseMob().User().BatchRegistry(context.Background(), []*user.RegisterReq{
		&user.RegisterReq{Username: "8888881", Password: "123456"},
		&user.RegisterReq{Username: "9999994", Password: "123456"}})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEaseMobUser_GetToken(t *testing.T) {
	data, err := newEaseMob().User().GetToken(context.Background(), &user.TokenReq{
		Username:       "888888341",
		GrantType:      "inherit",
		AutoCreateUser: false,
		Ttl:            utils.PointerAny(86400),
	})
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

	data, err = newEaseMob().User().GetToken(context.Background(), &user.TokenReq{
		Username:       "888888341",
		GrantType:      "inherit",
		AutoCreateUser: false,
		Ttl:            utils.PointerAny(86400),
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
