/**
 * @Author: Ne-21
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2022/1/12 20:59
 */

package controller

import (
	"fmt"
	"github.com/Olixn/Question-Bank-API/common"
	"github.com/Olixn/Question-Bank-API/common/utils"
	"github.com/Olixn/Question-Bank-API/database"
	"github.com/Olixn/Question-Bank-API/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var topicDao *model.TopicDao

func InitTopicDao() {
	var AppKvDbPool = database.AppKvDbPool
	var AppDb = database.AppDb

	topicDao = &model.TopicDao{
		KvDbPool: AppKvDbPool,
		Db:       AppDb,
	}
}

func GetNotice(c *gin.Context) {
	jsonData := utils.ReadJSONFile("././notice.json")
	c.JSON(http.StatusOK, jsonData)

}

func GetAnswer(c *gin.Context) {
	question, ok := c.GetQuery("q")
	fmt.Println(question)
	if !ok || question == "" {
		c.JSON(http.StatusOK, model.Response{
			Status: 0,
			Msg:    common.POST_QUESTION_EMPTY,
		})
		return
	}

	questionHASH := utils.MD5(question)
	answer, err := topicDao.Get(questionHASH)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Status: 0,
			Msg:    err.Error(),
		})
		return
	}
	var resData = model.ResData{
		Success: "true",
		Answer:  answer,
		Data:    answer,
	}

	c.JSON(http.StatusOK, model.Response{
		Status: 1,
		Msg:    "ok",
		Data:   resData,
	})

}
