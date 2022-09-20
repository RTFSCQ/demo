package main

import (
	"awesomeProject/pkg/mod/github.com/wonderivan/logger@v1.0.0"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	// 配置logger，如果不配置时默认为控制台输出，等级为DEBG
	logger.SetLogger(`{"Console": {"level": "DEBG"}`)
	// 配置说明见下文

	// 设置完成后，即可在控制台和日志文件app.log中看到如下输出
	logger.Trace("this is Trace")
	logger.Debug("this is Debug")
	logger.Info("this is Info")
	logger.Warn("this is Warn")
	logger.Error("this is Error")
	logger.Crit("this is Critical")
	logger.Alert("this is Alert")
	logger.Emer("this is Emergency")

}
