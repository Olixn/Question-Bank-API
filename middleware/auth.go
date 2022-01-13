/**
 * @Author: Ne-21
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2022/1/13 10:45
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func StartTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 处理下面的函数
		// c.Abort() //阻止执行
		stop := time.Since(start)
		fmt.Println(stop)
	}
}
