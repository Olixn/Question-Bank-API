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

	var TopicChan = make(chan map[int]*model.Topic, 100)

	topicDao = &model.TopicDao{
		KvDbPool:  AppKvDbPool,
		Db:        AppDb,
		TopicChan: TopicChan,
	}

	go topicDao.Write2Db()
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
			Msg:    common.ANSWER_NOT_FOUND,
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

func Update(c *gin.Context) {
	var question = c.PostForm("q")
	var answer = c.PostForm("a")
	var ip = c.ClientIP()

	if question == "" && answer == "" {
		c.JSON(http.StatusOK, model.Response{
			Status: 0,
			Msg:    common.POST_QUESTION_EMPTY,
		})
		return
	}

	var topic model.Topic
	topic.Question = question
	topic.Answer = answer
	topic.Ip = ip
	topic.Hash = utils.MD5(question)
	opType, _ := topicDao.SetAndUpdate(&topic)
	if opType == 1 {
		c.JSON(http.StatusOK, model.Response{
			Status: 1,
			Msg:    common.QUESTION_ADD,
		})
	} else if opType == 0 {
		c.JSON(http.StatusOK, model.Response{
			Status: 1,
			Msg:    common.QUESTION_UPDATE,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Status: 0,
			Msg:    common.UNKONW_ERROR,
		})
	}

}
