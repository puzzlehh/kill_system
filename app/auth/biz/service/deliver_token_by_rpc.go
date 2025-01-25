package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/auth"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	resp = &auth.DeliveryResp{
		Token: "123456",
	}
	time.Sleep(11 * time.Second)
	klog.Info("DeliverTokenReq:%+v", req)
	return resp, nil
}
