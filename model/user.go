package model

import "time"

type User struct {
	Id         int64
	Username   string     //昵称
	Account    string     //登录帐号
	Password   string     //密码
	PhoneNum   string     //电话号
	Status     int        //状态1:　系统初始化　2:用户已经完善信息 3:封号
	CreateTime time.Time//注册时间
}
