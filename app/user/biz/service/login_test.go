package service

import (
	"context"
	user "github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/user"
	"testing"
)

// 这里要怎么初始化数据连接？
func TestLogin_Run(t *testing.T) {
	// 这个ctx是空白的吗？
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "123",
		Password: "test",
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
