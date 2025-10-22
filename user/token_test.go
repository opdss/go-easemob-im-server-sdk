package user_test

import (
	"context"
	"fmt"
	"github.com/opdss/go-easemob-im-server-sdk/user"
	"github.com/opdss/go-easemob-im-server-sdk/utils"
	"testing"
)

func TestEasemobUser_Register(t *testing.T) {
	data, err := newEasemob().User().Register(context.Background(), &user.RegisterReq{Username: "8888881", Password: "123456"})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobUser_BatchRegistry(t *testing.T) {
	data, err := newEasemob().User().BatchRegistry(context.Background(), []*user.RegisterReq{
		&user.RegisterReq{Username: "8888881", Password: "123456"},
		&user.RegisterReq{Username: "9999994", Password: "123456"}})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}

func TestEasemobUser_GetToken(t *testing.T) {
	data, err := newEasemob().User().GetToken(context.Background(), &user.TokenReq{
		Username:       "888888341",
		GrantType:      "inherit",
		AutoCreateUser: false,
		Ttl:            utils.ToPointer(86400),
	})
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)

	data, err = newEasemob().User().GetToken(context.Background(), &user.TokenReq{
		Username:       "888888341",
		GrantType:      "inherit",
		AutoCreateUser: false,
		Ttl:            utils.ToPointer(86400),
	})

	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", err)
}
