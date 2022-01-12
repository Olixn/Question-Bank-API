/**
 * @Author: Ne-21
 * @Description: 主程序
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/1/12 16:20
 */

package main

import (
	"github.com/Olixn/Question-Bank-API/config"
	"github.com/Olixn/Question-Bank-API/route"
	"github.com/gin-gonic/gin"
)

var Config *config.Config

func init() {
	Config, _ = config.Init("config.yaml")
}

func main() {
	r := gin.Default()

	// 初始化路由
	route.Init(r)

	r.Run(":" + Config.Server.Port)

}
