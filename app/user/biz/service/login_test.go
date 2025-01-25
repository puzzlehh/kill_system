package service

import (
	"context"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/user"
	userRpc "github.com/puzzlehh/kill_system/rpc_gen/rpc/user"

	"testing"
)

// 这里要怎么初始化数据连接？
func TestLogin_Run(t *testing.T) {
	// 这个ctx是空白的吗？
	ctx := context.Background()
	// init req and assert value
	r, err := consul.NewConsulResolver("10.21.32.14:8500")
	c, err := userRpc.NewRPCClient("user", client.WithResolver(r))

	req := &user.LoginReq{
		Email:    "123",
		Password: "test",
	}

	resp, err := c.Login(ctx, req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
