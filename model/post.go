package model

import (
	"fmt"
	"github.com/frustra/bbcode"
	"log"
)

type Post struct {
	tbname string
}

type dzPost struct {
	pid,
	tid,
	fid,
	authorid,
	first,
	message,
	dateline string
}

type hyPost struct {
	id,
	tid,
	fid,
	uid,
	isthread,
	content,
	atime string
}

func (p *Post) Init() (err error) {
	p.tbname = "hy_post"
	return p.ToConvert()
}

func (p *Post) ToConvert() (err error) {
	log.Println("正在转换 " + p.tbname + " ...")

	err = Truncate(p.tbname)
	if err != nil {
		return
	}

	dzSqlStr := "SELECT pid, tid, fid, authorid, first, message, dateline FROM `cdb_posts`"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, tid, fid, uid, isthread, content, atime) VALUES (?, ?, ?, ?, ?, ?, ?)", p.tbname)

	data, err := DiscuzDbTx.Query(dzSqlStr)
	if err != nil {
		log.Println("Dz post 查询失败: " + dzSqlStr)
		log.Println(err)
		return
	}

	stmt, err := HybbsDbTx.Prepare(hySqlStr)
	if err != nil {
		log.Println("Hy post 预加载失败: " + hySqlStr)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyPost
	var content string

	for data.Next() {
		d1 := new(dzPost)
		err = data.Scan(&d1.pid, &d1.tid, &d1.fid, &d1.authorid, &d1.first, &d1.message, &d1.dateline)
		if err != nil {
			log.Println("Dz post 扫描取值失败")
			log.Println(err)
			return
		}

		//bbcode 转 html
		compiler := bbcode.NewCompiler(true, true)
		content = compiler.Compile(d1.message)

		hydata := hyPost{
			d1.pid,
			d1.tid,
			d1.fid,
			d1.authorid,
			d1.first,
			content,
			d1.dateline,
		}

		dataArr = append(dataArr, hydata)
	}
	defer DiscuzDbTx.Rollback()
	DiscuzDbTx.Commit()

	for _, value := range dataArr {
		_, err = stmt.Exec(value.id, value.tid, value.fid, value.uid, value.isthread, value.content, value.atime)
		if err != nil {
			return
		}

		stat++
	}
	defer HybbsDbTx.Rollback()
	HybbsDbTx.Commit()

	if err == nil {
		log.Printf("%s 转换成功, 总共插入 %d 条数据", p.tbname, stat)
	} else {
		log.Printf("%s 转换失败", p.tbname)
	}
	return
}
