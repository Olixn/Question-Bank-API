/**
 * @Author: Ne-21
 * @Description:
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2022/1/12 21:01
 */

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"os"
)

func ReadJSONFile(filepath string) (jsonData map[string]interface{}) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(f)

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return nil
	}
	jsonData = v.(map[string]interface{})
	return
}

func Store2Redis(conn redis.Conn, h string, answer string) {
	_, err := conn.Do("Set", h, answer)
	if err != nil {
		fmt.Println("Store2Redis Err = ", err)
		return
	}
}
