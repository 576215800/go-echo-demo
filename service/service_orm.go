package service

import (
	"database/sql"
	logger "echodemo/logs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	db_driver = "root:root@tcp(127.0.0.1:3306)/blogs?charset=utf8"
)

var (
	db *sql.DB
)

/**
初始化数据库连接
测试执行查询t_user总数
*/
func start() {
	var err error
	logger.Info("start connection databases....")
	db, err = sql.Open("mysql", db_driver)
	if err != nil {
		logger.Error("数据库打开出现了问题.....")
		panic(err)
	}
	//测试数据库查询
	err = db.Ping()
	if err!=nil{
		log.Fatal("数据库连接出现了问题....")
		panic(err)
	}
}
