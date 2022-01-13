/**
 * @Author: Ne-21
 * @Description: 路由
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2022/1/12 17:48
 */

package route

import (
	"github.com/Olixn/Question-Bank-API/controller"
	"github.com/Olixn/Question-Bank-API/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://script.521daigua.cn")
	})

	apiGroup := r.Group("/v1")
	apiGroup.Use(middleware.StartTime())
	{
		apiGroup.GET("/api/notice", controller.GetNotice)
		apiGroup.GET("/api/answer", controller.GetAnswer)
	}

}
