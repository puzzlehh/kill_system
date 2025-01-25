package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/puzzlehh/kill_system/app/user/conf"
	"github.com/puzzlehh/kill_system/common/clientsuite"
	"github.com/puzzlehh/kill_system/rpc_gen/kitex_gen/auth/authservice"
	"sync"
)

var (
	AuthClient   authservice.Client
	once         sync.Once
	err          error
	registryAddr string
	serviceName  string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		initClient("auth")
	})
}

func initClient(serviceName string) {
	AuthClient, err = authservice.NewClient(serviceName, commonSuite)
	//todo: 错误处理？
	//checkoututils.MustHandleError(err)
}
