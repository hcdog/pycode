package auth

import (
	"context"

	"github.com/hcdog/pycode/gomall/app/frontend/biz/service"
	"github.com/hcdog/pycode/gomall/app/frontend/biz/utils"
	auth "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusOK, []byte("/"))
}

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq//隐性初始化
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}	
	// resp := &common.Empty{}
	redirect, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte(redirect))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	_, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
