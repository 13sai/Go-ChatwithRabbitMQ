package util

import (
	"fmt"
	"bytes"
	"log"
	"time"	
	"crypto/md5"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 合并字符串
func StrCombine(str ...string) string {
	var bt bytes.Buffer
	for _, arg := range str {
		bt.WriteString(arg)
    }
	//获得拼接后的字符串
	return bt.String()
}

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)));
}
