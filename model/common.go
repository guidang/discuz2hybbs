package model

import (
	"fmt"
	"database/sql"
	"log"
)

/**
 清空table
 */
func Truncate(db *sql.DB, tb string) (err error) {
	sqlStr := fmt.Sprintf("TRUNCATE TABLE %s", tb)
	_, err = db.Exec(sqlStr);

	if err != nil {
		log.Println(err)
		log.Printf("清空表 %s 失败", tb)
	} else {
		log.Printf("清空表 %s 成功", tb)
	}
	return
}
