package filter

import (
	"echodemo/model"
	"github.com/labstack/echo"
)

var (
	User *model.User
)
/*
	此方法为公共拦截器
	做是否登录　权限校验等操作
	给全局变量设值可以在Controller中获取到
 */
func OpenTracing() echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error{
			//从请求头中获取请求参数
			//token:=c.Request().Header.Get("t")
			//uid:=c.Request().Header.Get("u")
			return next(c)
		}
	}
}
