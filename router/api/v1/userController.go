package v1

import (
	logger "echodemo/logs"
	"echodemo/model"
	"echodemo/service"
	"github.com/labstack/echo"
	"net/http"
)

/*
	注册用户信息
*/
func InsertUser(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	restRes := &RestResponse{}
	if err != nil {
		logger.Error("解析前端请求出错")
		restRes.Status = http.StatusBadRequest
		restRes.Msg = "请求参数错误...."
		return c.JSON(http.StatusBadRequest, restRes)
	}
	id, err := service.InsertUser(user)
	if err != nil {
		restRes.Status = http.StatusOK
		restRes.Msg = "用户注册失败..."
		restRes.Code = 4000
		restRes.Data = id
	}
	restRes.Status = http.StatusOK
	restRes.Msg = "用户注册成功..."
	restRes.Code = 2000
	restRes.Data = id
	return c.JSON(http.StatusOK, user)
}

/*
	查询用户信息
*/
func GetUser(c echo.Context) error {
	//得到url path上的参数
	id := c.Param("account")
	logger.Debug(id)
	return c.JSON(http.StatusOK, "查询用户...")
}

/*
 删除用户信息
*/
func DeleteUser(c echo.Context) error {
	id := c.Param("account")
	logger.Debug(id)
	return c.JSON(http.StatusOK, "删除用户...")
}

/*
修改用户密码
*/
func UpdateUser(c echo.Context) error {
	account := c.QueryParam("account")
	password := c.QueryParam("password")
	logger.Debug(account, password)
	return c.JSON(http.StatusOK, "修改用户信息...")
}
