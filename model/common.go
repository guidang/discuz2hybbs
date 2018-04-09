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

func SetConvertLog(str string, code int) {
	fmt.Println(str, code)
}
