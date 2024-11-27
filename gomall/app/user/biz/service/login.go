package service

import (
	"context"
	"errors"

	"github.com/hcdog/pycode/gomall/app/user/biz/dal/mysql"
	"github.com/hcdog/pycode/gomall/app/user/biz/model"
	user "github.com/hcdog/pycode/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Password == "" {
		err = errors.New("请输入密码")
		return
	}
	//得到数据库里的数据
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		return
	}
	return &user.LoginResp{UserId: int32(userRow.ID)}, nil

}
