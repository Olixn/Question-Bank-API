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
	KvDbPool *redis.Pool
	Db       *gorm.DB
}

func (t *TopicDao) Get(h string) (answer string, err error) {
	kvConn := t.KvDbPool.Get()
	defer kvConn.Close()

	answer, err = redis.String(kvConn.Do("Get", h))
	if err == nil && answer != "" {
		return
	}

	var dbRes Topic
	res := t.Db.Table("topic").Where("hash = ?", h).Select("answer").Find(&dbRes)
	if res.RowsAffected > 0 {
		err = nil
		answer = dbRes.Answer
		utils.Store2Redis(kvConn, h, answer)
		return
	}
	return
}
