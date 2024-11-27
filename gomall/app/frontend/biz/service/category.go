package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/category"
	// common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string] any, err error) {
	//商品列表
	p, _ := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
