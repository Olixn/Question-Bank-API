/**
 * @Author: Ne-21
 * @Description:
 * @File:  md5
 * @Version: 1.0.0
 * @Date: 2022/1/13 12:53
 */

package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) (s string) {
	data := []byte(str)
	s = fmt.Sprintf("%x", md5.Sum(data))
	return
}
