package rpc

import (
	"sync"

	// "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	checkoututils "github.com/hcdog/pycode/gomall/app/checkout/utils"
	// "github.com/hcdog/pycode/gomall/app/cart/conf"
	// cartutils "github.com/hcdog/pycode/gomall/app/cart/utils"
	// "github.com/hcdog/pycode/gomall/app/checkout/conf"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	consul "github.com/kitex-contrib/registry-consul"
)


var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initCartClient()
		initPaymentClient()

	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	checkoututils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	checkoututils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.0.1:8500")
	checkoututils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}
