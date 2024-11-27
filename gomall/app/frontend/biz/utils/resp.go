package utils

import (
	// "app/frontend/middleware"
	"context"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/hertz-contrib/sessions"
	"github.com/hcdog/pycode/gomall/app/frontend/utils"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/cart"
	frontendutils "github.com/hcdog/pycode/gomall/app/frontend/utils"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any{
	var cartNum int
	use_id := frontendutils.GetUserIdFromCtx(ctx)
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: use_id})
	if cartResp != nil && cartResp.Cart != nil {
		cartNum = len(cartResp.Cart.Items)
	}
	content["user_id"] = ctx.Value(utils.UserIdKey)
	content["cart_num"] = cartNum
	return content
}