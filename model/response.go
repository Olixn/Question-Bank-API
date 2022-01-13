/**
 * @Author: Ne-21
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:43
 */

package model

type Response struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   ResData `json:"data"`
}

type ResData struct {
	Success string `json:"success,omitempty"`
	Answer  string `json:"answer,omitempty"`
	Data    string `json:"data,omitempty"`
}
