/**
 * @Author: Ne-21
 * @Description:
 * @File:  topic
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:02
 */

package model

type Topic struct {
	Id        int    `json:"id,omitempty"`
	Hash      string `json:"hash,omitempty"`
	Question  string `json:"question,omitempty"`
	Answer    string `json:"answer,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Course    string `json:"course,omitempty"`
	ChapterId string `json:"chapterId,omitempty"`
}
