/**
 * @Author: Ne-21
 * @Description:
 * @File:  topicDao
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:09
 */

package model

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type TopicDao struct {
	KvDbPool *redis.Pool
	Db       *gorm.DB
}

func (t *TopicDao) Get(h string) (answer string, err error) {
	kvConn := t.KvDbPool.Get()

	answer, err = redis.String(kvConn.Do("Get", h))
	if err == nil {
		return
	}

	return
}
