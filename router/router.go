package router

import (
	"echodemo/logs"
	ourmiddleware "echodemo/middleware"
	v1 "echodemo/router/api/v1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

/*
	echo初始化设置
*/
func Start() {
	logger.Debug("start router....")
	//初始化echo
	e = echo.New()
	//设置静态资源文件路径
	e.Static("/", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ourmiddleware.Count)
	initRouter()
	//自定义中间件拦截器
	//
	e.Logger.Fatal(e.Start(":1323"))
}

/*
	初始化路由
*/
func initRouter() {
	//用户组定义中间件
	userGroup := e.Group("/api/users")
	userGroup.GET("/user/:id", v1.GetUser)
	userGroup.POST("/user", v1.InsertUser)
	userGroup.DELETE("/user/:id", v1.DeleteUser)
	userGroup.PUT("/user/:id", v1.UpdateUser)
}
