package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/cart"
	rpccart "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/cart"
	common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"
	frontendutils "github.com/hcdog/pycode/gomall/app/frontend/utils"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	return
}
