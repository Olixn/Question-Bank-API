/**
 * @Author: Ne-21
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2022/1/12 20:59
 */

package controller

import (
	"github.com/Olixn/Question-Bank-API/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCxNotice(c *gin.Context) {
	jsonData := utils.ReadJSONFile("././notice.json")
	c.JSON(http.StatusOK, jsonData)
}
