package middleware

import (
	"context"
	"fmt"

	"github.com/hcdog/pycode/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		// 使用 Next 使得路由匹配的函数执行
		c.Next(ctx)//执行下一个操作
		// 在 Next 中的函数执行之后打印日志
	}
}

func Auth() app.HandlerFunc{
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				next = fmt.Sprintf("%s?next=%s", next, ref)
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
		// 使用 Next 使得路由匹配的函数执行
		c.Next(ctx)//执行下一个操作
		// 在 Next 中的函数执行之后打印日志
	}


}