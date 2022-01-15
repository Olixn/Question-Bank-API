/**
 * @Author: Ne-21
 * @Description:
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2022/1/13 11:17
 */

package database

import (
	"fmt"
	"github.com/Olixn/Question-Bank-API/config"
	"github.com/Olixn/Question-Bank-API/model"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

var AppDb *gorm.DB

func InitMySql() {
	var err error
	var AppConfig = config.AppConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.MySql.UserName,
		AppConfig.MySql.PassWord,
		AppConfig.MySql.Address,
		AppConfig.MySql.Port,
		AppConfig.MySql.Dbname)
	AppDb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

	err = AppDb.AutoMigrate(&model.Topic{})
	if err != nil {
		fmt.Println("数据表初始化失败")
		panic(err)
		return
	}
}
