package controller

import "C"
import (
	"github.com/labstack/echo"
	"net/http"
)

/*
	注册用户信息
*/
func InsertUser(c echo.Context) error {

	return c.JSON(http.StatusOK, "注册用户.....")
}

/*
	查询用户信息
*/
func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "查询用户...")
}

/*
 删除用户信息
*/
func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "删除用户...")
}
/*
　修改用户信息
 */
func UpdateUser(c echo.Context) error{
	return c.JSON(http.StatusOK,"修改用户信息...")
}