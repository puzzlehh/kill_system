package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/puzzlehh/kill_system/app/user/infra/rpc"
	"github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/auth"
	user "github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/user"
	"strconv"
)

// 这里是干啥的
type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	klog.Infof("LoginReq:%+v", req)

	//if mysql.DB == nil {
	//	klog.Error("Database connection is nil")
	//	return nil, errors.New("database connection is nil")
	//}
	//
	//userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	//
	//if err != nil {
	//	return
	//}
	//if userRow == nil {
	//	return
	//}
	//if req.Password != userRow.Password {
	//	return
	//}

	//可以抽离
	result, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{})
	if err != nil {
		klog.CtxErrorf(s.ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	klog.Infof("result:%+v", result)

	v, _ := strconv.Atoi(result.Token)
	resp = &user.LoginResp{
		UserId: int32(v),
	}

	return resp, nil
}
