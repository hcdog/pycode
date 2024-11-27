package service

import (
	"context"

	common "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo register
	userRspc ,err := rpc.UserClient.Register(h.Context,&user.RegisterReq{
		Email: req.Email,
		Password: req.Password,
		ConfirmPassword: req.ConfirmPassword,	
	})
	if err != nil{
		return nil,err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userRspc.UserId)
	session.Save()

	return resp,nil
}
