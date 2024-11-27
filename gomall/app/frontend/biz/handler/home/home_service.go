package home

import (

	"context"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/utils"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"
	// home "app/frontend/hertz_gen/frontend/home"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	var resp map[string]any
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err = service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	
	c.HTML(consts.StatusOK, "home.tmpl", utils.WarpResponse(ctx, c, resp))

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
