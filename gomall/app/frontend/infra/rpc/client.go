package rpc

import (
	"sync"

	// "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	// "github.com/cloudwego/kitex/pkg/circuitbreak"
	// "github.com/cloudwego/kitex/pkg/rpcinfo"
	frontendutils "github.com/hcdog/pycode/gomall/app/frontend/utils"

	// "github.com/hcdog/pycode/gomall/app/frontend/conf"
	consul "github.com/kitex-contrib/registry-consul"

	// "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client

	once           sync.Once
	// err            error
	// registryAddr   string
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})//这里有问题，要是初始化失败，所有的都要再来一次
}

func initUserClient() {
	// 用户服务注册
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	frontendutils.MustHandleError(err)
	//加入新用户
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendutils.MustHandleError(err)
}
func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	frontendutils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}
func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	frontendutils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendutils.MustHandleError(err)
}
