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
	AppDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
}
