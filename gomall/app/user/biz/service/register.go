package service

import (
	"context"
	// "errors"

	"github.com/hcdog/pycode/gomall/app/user/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/user/biz/model"
	user "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)
type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// if req.Email == "" {
	// 	err = errors.New("请输入email")
	// 	return
	// }
	//接受输入，一个结构体，包含了我设置的字段，邮箱，密码和确认密码
	// if req.Password != req.ConfirmPassword {
	// 	err = errors.New("密码不一")
	// 	return
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)//hash密码
	if err != nil {
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	//创建user.ID
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil//返回注册信息

}
