package model

import (
	"database/sql"
	"log"
	"fmt"
	"regexp"
)

var (
	tbname string = "hy_forum"
	tbname2 string = "hy_forum_group"
)

type Forum struct {
	Hdb *sql.DB
	Ddb *sql.DB
}

type dzForum struct {
	fid,
	fup,
	types,
	name,
	threads,
	posts,
	desc string
}

type hyForum struct {
	id,
	fid,
	fgid,
	name,
	threads,
	posts,
	desc string
}

func (f * Forum) Init(hdb, ddb *sql.DB) (err error){
	f.Hdb = hdb
	f.Ddb = ddb

	return f.ToConvert()
}

func (f *Forum) ToConvert() (err error) {
	log.Println("forum ToConvert")

	err = Truncate(f.Hdb, tbname)
	if err != nil {
		return
	}

	dzSqlStr := "SELECT f.fid, f.fup, f.type, f.name, f.threads, f.posts, z.description  FROM `cdb_forums` f LEFT JOIN `cdb_forumfields` z ON f.fid = z.fid"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, fid, fgid, name, threads, posts, html, name2, forumg, json, color, background) VALUES (?, ?, ?, ?, ?, ?, ?, '', '', '', '', '')", tbname)

	data, err := f.Ddb.Query(dzSqlStr)
	if err != nil {
		log.Println("Dz forum 查询失败: " + dzSqlStr)
		log.Println(err)
		return
	}

	stmt, err := f.Hdb.Prepare(hySqlStr)
	if err != nil {
		log.Println("Hy forum 预加载失败: " + hySqlStr)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyForum
	cateGroup := make(map[string]string)
	groupMap := make(map[string]string)

	for data.Next() {
		d1 := new(dzForum)
		err = data.Scan(&d1.fid, &d1.fup, &d1.types, &d1.name, &d1.threads, &d1.posts, &d1.desc);
		if err != nil {
			log.Println("Dz forum 扫描取值失败")
			log.Println(err)
			return
		}

		var fid, fgid, name string
		if d1.types == "forum" {
			fid = "-1"
			fgid = d1.fup
		} else {
			fid = d1.fup
			fgid = "0"
		}

		name = filterName(d1.name, 12)
		log.Println(name)

		if d1.types == "group" {
			groupMap[d1.fid] = filterName(d1.name, 32)
			continue
		}

		if d1.types == "forum" {
			cateGroup[d1.fid] = d1.fup
		}

		hydata := hyForum{
			d1.fid,
			fid,
			fgid,
			name,
			d1.threads,
			d1.posts,
			d1.desc,
		}

		dataArr = append(dataArr, hydata)
	}

	//log.Println(groupMap)
	//分组表
	stmt2, err := f.Hdb.Prepare(fmt.Sprintf("INSERT INTO %s (id, name) VALUES (?, ?)", tbname2))
	if err != nil {
		log.Println(tbname2 + " 预加载失败: ")
		log.Println(err)
		return
	}
	err = Truncate(f.Hdb, tbname2)
	if err != nil {
		return
	}
	for index, value := range groupMap {
		fmt.Printf("arr[%d]=%s \n", index, value)
		_, err = stmt2.Exec(index, value)
		if err != nil {
			log.Println(tbname2 + " 导入失败")
			return
		}
	}

	//log.Println(cateGroup)

	//log.Println(dataArr)
	//分类表
	for _, value := range dataArr {
		//fmt.Printf("arr[%d]=%s \n", index, value)
		if value.fgid == "0" {
			value.fgid = cateGroup[value.fid]
		}
		_, err = stmt.Exec(value.id, value.fid, value.fgid, value.name, value.threads, value.posts, value.desc)
		if err != nil {
			return
		}

		stat++
	}

	if err == nil {
		log.Printf("%s 转换成功, 总共插入 %d 条数据", tbname, stat)
	} else {
		log.Printf("%s 转换失败", tbname)
	}
	return
}

func filterName(str string, lenght int) string {
	log.Println("filterName")
	var match = regexp.MustCompile(`-\.-(.*)`)
	ret := match.FindStringSubmatch(str)
	if len(ret) > 1 {
		str = ret[1]
	}

	nameStr := []rune(str)
	lth := len(nameStr)
	if lenght > lth {
		lenght = lth
	}
	return string(nameStr[: lenght])
}