/**
 * @Author: Ne-21
 * @Description:
 * @File:  topic
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:02
 */

package model

type Topic struct {
	Id        int
	Hash      string `json:"hash"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	Ip        string `json:"ip"`
	Course    string `json:"course"`
	ChapterId string `json:"chapterId"`
}
