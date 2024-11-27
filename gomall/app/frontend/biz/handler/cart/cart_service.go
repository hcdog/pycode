package cart

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/service"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/utils"
	cart "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_ , err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}

	c.Redirect(consts.StatusFound, []byte("/cart"))
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}
