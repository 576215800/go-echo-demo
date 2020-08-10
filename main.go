package main

import (
	"echodemo/router"
	"echodemo/logs"
)

func main() {
	//初始化日志组件
	initLog()
	router.Start()

}
func initLog() {
	var cfg = logger.LogConfig{Type: "date", Root: "./logs", Maxdays: "7"}
	logger.Init(cfg)
}
