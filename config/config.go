/**
 * @Author: Ne-21
 * @Description: 配置信息
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/1/12 16:28
 */

package config

import (
	"errors"
	"fmt"
	"github.com/Olixn/Question-Bank-API/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var AppConfig *Config

type Config struct {
	Server *Server
	MySql  *MySql
	Redis  *Redis
}

type Server struct {
	Port string `yaml:"port"`
}

type MySql struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

type Redis struct {
	Address     string `yaml:"address"`
	Port        string `yaml:"port"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

func Init(filename string) {
	AppConfig = &Config{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.New(common.CONFIG_INIT_ERROR)
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	fmt.Println(string(content))
	if yaml.Unmarshal(content, &AppConfig) != nil {
		err = errors.New(common.CONFIG_INIT_ERROR)
		log.Fatalf("解析config.yaml出错: %v", err)
	}
}
