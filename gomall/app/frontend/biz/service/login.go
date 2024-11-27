package service

import (
	"context"

	auth "github.com/hcdog/pycode/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
	"github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user"
	// "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user/userservice"

	// common "app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/hcdog/pycode/gomall/app/frontend/infra/rpc"
	"github.com/hertz-contrib/sessions"
	// frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// resp ,err := rpc.UserClient.Login(h.Context,&user.LoginReq{
	// 				Email: req.Email,
	// 				Password: req.Password,	
	// 			})
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil{
		return
	}
	//user svc api
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	session.Save()
	// frontendutils.MustHandleError(err)
	redirect = "/"
	// // if frontendutils.ValidateNext(req.Next) {
	// // 	redirect = req.Next
	// // }
	
	if redirect != "" {
		redirect = req.Next
	}


	return redirect,err
}
