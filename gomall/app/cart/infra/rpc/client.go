package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	// "github.com/hcdog/pycode/gomall/app/cart/conf"
	cartutils "github.com/hcdog/pycode/gomall/app/cart/utils"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	// err           error
	// registryAddr  string
	// serviceName   string
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	cartutils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}

