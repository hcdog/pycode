// 鉴权相关代码
package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterMiddleware(h *server.Hertz) {
	h.Use(GlobalAuth())//use函数，写入中间件
}