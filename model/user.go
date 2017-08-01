package model

import (
	"fmt"
	"log"
)

type User struct {
	adminid,
	tbname string
}

type dzUser struct {
	uid,
	username,
	password,
	email,
	threads,
	posts,
	regdate,
	credits,
	lastvisit,
	salt string
}

type hyUser struct {
	id,
	user,
	pass,
	email,
	threads,
	posts,
	atime,
	credits,
	ctime,
	salt string
}

func (u *User) Init() (err error) {
	u.tbname = "hy_user"
	err = u.ToConvert()
	if err == nil {
		err = u.setManager()
	}
	return
}

func (u *User) ToConvert() (err error) {
	log.Println("user ToConvert")

	err = Truncate(HybbsDb, u.tbname)
	if err != nil {
		return
	}

	//dzSqlStr := "SELECT uid, username, password, email, threads, posts, regdate, credits, lastvisit FROM `cdb_members`"
	dzSqlStr := "SELECT m.uid, m.username, c.password, m.email, m.threads, m.posts, m.regdate, m.credits, m.lastvisit, c.salt FROM `cdb_members` m LEFT JOIN `cdb_uc_members` c ON c.uid = m.uid WHERE c.salt IS NOT NULL"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, user, pass, email, threads, posts, atime, credits, ctime, salt, `group`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 2)", u.tbname)

	data, err := DiscuzDb.Query(dzSqlStr)
	if err != nil {
		log.Println("Dz user 查询失败: " + dzSqlStr)
		log.Println(err)
		return
	}

	stmt, err := HybbsDb.Prepare(hySqlStr)
	if err != nil {
		log.Println("Hy user 预加载失败: " + hySqlStr)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyUser

	for data.Next() {
		d1 := new(dzUser)
		err = data.Scan(&d1.uid, &d1.username, &d1.password, &d1.email, &d1.threads, &d1.posts, &d1.regdate, &d1.credits, &d1.lastvisit, &d1.salt)
		if err != nil {
			log.Println("Dz user 扫描取值失败")
			log.Println(err)
			return
		}

		hydata := hyUser{
			d1.uid,
			d1.username,
			d1.password,
			d1.email,
			d1.threads,
			d1.posts,
			d1.regdate,
			d1.credits,
			d1.lastvisit,
			d1.salt,
		}

		dataArr = append(dataArr, hydata)
	}

	for _, value := range dataArr {
		_, err = stmt.Exec(value.id, value.user, value.pass, value.email, value.threads, value.posts, value.atime, value.credits, value.ctime, value.salt)
		if err != nil {
			return
		}

		stat++
	}

	if err == nil {
		log.Printf("%s 转换成功, 总共插入 %d 条数据", u.tbname, stat)
	} else {
		log.Printf("%s 转换失败", u.tbname)
	}
	return
}

func (u *User) setManager() (err error) {
	hySqlStr := fmt.Sprintf("UPDATE %s SET `group` = 1 WHERE id = ?", u.tbname)

	if u.adminid == "" {
		u.adminid = "1"
	}

	_, err = HybbsDb.Exec(hySqlStr, u.adminid)
	if err != nil {
		log.Println("Dz user 设置管理员失败")
		log.Println(err)
		return
	} else {
		log.Printf("\r\nDz user 设置管理员为 %s 成功\r\n", u.adminid)
	}
	return
}
