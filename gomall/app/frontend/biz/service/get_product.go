package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	rpcproduct "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product"

	product "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//商品详情
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"item": p.Product,
	}, nil

}
