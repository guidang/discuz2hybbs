package model

import (
	"fmt"
	"log"
)

/**
清空table
*/
func Truncate(tb string) (err error) {
	sqlStr := fmt.Sprintf("TRUNCATE TABLE %s", tb)

	_, err = HybbsDb.Exec(sqlStr)
	if err != nil {
		log.Println(err)
		log.Printf("清空表 %s 失败", tb)
	} else {
		log.Printf("清空表 %s 成功", tb)
	}
	return
}

func SetConvertLog(msg string, code int) {
	if code != 2 && msg != "" {
		log.Println(msg)
	}
	msg = Te.Text() + "\r\n" + msg
	Te.SetText(msg)
}
