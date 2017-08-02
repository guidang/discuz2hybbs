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
	SetConvertLog("正在转换 "+u.tbname+" ...", 0)

	err = Truncate(u.tbname)
	if err != nil {
		return
	}

	//dzSqlStr := "SELECT uid, username, password, email, threads, posts, regdate, credits, lastvisit FROM `cdb_members`"
	dzSqlStr := "SELECT m.uid, m.username, c.password, m.email, m.threads, m.posts, m.regdate, m.credits, m.lastvisit, c.salt FROM `cdb_members` m LEFT JOIN `cdb_uc_members` c ON c.uid = m.uid WHERE c.salt IS NOT NULL"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, user, pass, email, threads, posts, atime, credits, ctime, salt, `group`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 2)", u.tbname)

	data, err := DiscuzDbTx.Query(dzSqlStr)
	if err != nil {
		SetConvertLog("Dz user 查询失败: "+dzSqlStr, -1)
		log.Println(err)
		return
	}

	stmt, err := HybbsDb.Prepare(hySqlStr)
	if err != nil {
		SetConvertLog("Hy user 预加载失败: "+hySqlStr, -1)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyUser

	for data.Next() {
		d1 := new(dzUser)
		err = data.Scan(&d1.uid, &d1.username, &d1.password, &d1.email, &d1.threads, &d1.posts, &d1.regdate, &d1.credits, &d1.lastvisit, &d1.salt)
		if err != nil {
			SetConvertLog("Dz user 扫描取值失败", -1)
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
	defer DiscuzDbTx.Rollback()
	DiscuzDbTx.Commit()

	for _, value := range dataArr {
		_, err = stmt.Exec(value.id, value.user, value.pass, value.email, value.threads, value.posts, value.atime, value.credits, value.ctime, value.salt)
		if err != nil {
			return
		}

		stat++
	}
	defer HybbsDbTx.Rollback()
	HybbsDbTx.Commit()

	if err == nil {
		msg := fmt.Sprintf("%s 转换成功, 总共插入 %d 条数据", u.tbname, stat)
		SetConvertLog(msg, 0)
	} else {
		msg := fmt.Sprintf("%s 转换失败")
		SetConvertLog(msg, -1)
		log.Println(err)
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
		SetConvertLog("Hy user 设置管理员失败", -1)
		log.Println(err)
		return
	} else {
		msg := fmt.Sprintf("Hy user 设置管理员为 %s 成功", u.adminid)
		SetConvertLog(msg, -1)
		log.Println(err)
	}
	return
}
