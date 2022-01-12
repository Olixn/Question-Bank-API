/**
 * @Author: Ne-21
 * @Description: 路由
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2022/1/12 17:48
 */

package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, nil)
	})
}
