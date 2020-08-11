package middleware

import (
	"fmt"
	"github.com/labstack/echo"
)
//记录访问量
var totalRequests = 0
//自定义中间件
func Count(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{
		totalRequests++
		c.Response().Header().Add("requests",fmt.Sprintf("%d", totalRequests))
		return next(c)
	}
}
