/**
 * @Author: Ne-21
 * @Description:
 * @File:  topicDao
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:09
 */

package model

import (
	"github.com/Olixn/Question-Bank-API/common/utils"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type TopicDao struct {
	KvDbPool  *redis.Pool
	Db        *gorm.DB
	TopicChan chan map[int]*Topic
}

func (t *TopicDao) Get(h string) (answer string, err error) {
	kvConn := t.KvDbPool.Get()
	defer kvConn.Close()

	answer, err = redis.String(kvConn.Do("Get", h))
	if err == nil && answer != "" {
		return
	}

	var dbRes Topic
	res := t.Db.Model(&Topic{}).Where("hash = ?", h).Select("answer").Find(&dbRes)
	if res.RowsAffected > 0 {
		err = nil
		answer = dbRes.Answer
		utils.Store2Redis(kvConn, h, answer)
		return
	}
	return
}

func (t *TopicDao) SetAndUpdate(topic *Topic) (opType int, err error) {
	var dbRes Topic
	res := t.Db.Model(&Topic{}).Where("hash = ?", topic.Hash).Find(&dbRes)
	if res.RowsAffected > 0 {
		dbRes.Answer = topic.Answer
		opType = 0
		t.TopicChan <- map[int]*Topic{opType: &dbRes}
	} else {
		opType = 1
		t.TopicChan <- map[int]*Topic{1: topic}
	}
	return
}

func (t *TopicDao) Write2Db() {
	for {
		topicMap := <-t.TopicChan
		for k, v := range topicMap {
			if k == 0 {
				t.Db.Model(&v).Updates(v)
			} else {
				t.Db.Model(&Topic{}).Create(v)
			}
		}
	}
}
