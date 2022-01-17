/**
 * @Author: Ne-21
 * @Description: ip流控限制中间件
 * @File:  ipRestrict
 * @Version: 1.0.0
 * @Date: 2022/1/16 15:19
 */

package middleware

import (
	"github.com/Olixn/Question-Bank-API/database"
	"github.com/Olixn/Question-Bank-API/model"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"strings"
)

func IPRestrict() gin.HandlerFunc {
	/*
		对求的IP进行频率限制
	*/
	return func(c *gin.Context) {
		var (
			count       = 10 // 频次数
			cycle       = 60 // 时间周期 单位（second）
			AppKvDbPool = database.AppKvDbPool
		)
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		conn := AppKvDbPool.Get()
		defer conn.Close()
		defer conn.Do("Select", 0)
		_, err := conn.Do("Select", 1)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Status: 403,
				Msg:    "Redis Server Error",
			})
			c.Abort() //终止请求，直接返回提示信息
			return
		}
		res, _ := redis.String(conn.Do("Get", ip))
		if res == "" {
			_, err := conn.Do("SetEx", ip, cycle, count) //conn.Do()用的是最多的，把命令行中的args一个个写进去
			if err != nil {
				c.JSON(http.StatusOK, model.Response{
					Status: 405,
					Msg:    "server error",
				})
				c.Abort() //终止请求，直接返回提示信息
				return
			}
		}
		if res == "1" {
			c.JSON(http.StatusOK, model.Response{
				Status: 405,
				Msg:    "触发流控限制，客官慢一些",
			})
			c.Abort() //终止请求，直接返回提示信息
			return
		}
		conn.Do("DECR", ip)
	}
}
