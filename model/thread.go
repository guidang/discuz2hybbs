package model

import (
	"fmt"
	"log"
)

type Thread struct {
	tbname string
}

type dzThread struct {
	tid,
	fid,
	authorid,
	pid,
	subject,
	dateline,
	lastpost,
	views,
	replies,
	attachment string
}

type hyThread struct {
	id,
	fid,
	uid,
	pid,
	title,
	atime,
	btime,
	views,
	posts,
	files string
}

func (t *Thread) Init() (err error) {
	t.tbname = "hy_thread"
	return t.ToConvert()
}

func (t *Thread) ToConvert() (err error) {
	log.Println("正在转换 " + t.tbname + " ...")

	err = Truncate(t.tbname)
	if err != nil {
		return
	}

	dzSqlStr := "SELECT t.tid, t.fid, t.authorid, p.pid, t.subject, t.dateline, t.lastpost, t.views, t.replies, t.attachment FROM `cdb_threads` t LEFT JOIN cdb_posts p ON p.tid = t.tid WHERE p.first = 1"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, fid, uid, pid, title, atime, btime, views, posts, files, summary, img) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, '', '')", t.tbname)

	data, err := DiscuzDbTx.Query(dzSqlStr)
	if err != nil {
		log.Println("Dz thread 查询失败: " + dzSqlStr)
		log.Println(err)
		return
	}

	stmt, err := HybbsDbTx.Prepare(hySqlStr)
	if err != nil {
		log.Println("Hy thread 预加载失败: " + hySqlStr)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyThread

	for data.Next() {
		d1 := new(dzThread)
		err = data.Scan(&d1.tid, &d1.fid, &d1.authorid, &d1.pid, &d1.subject, &d1.dateline, &d1.lastpost, &d1.views, &d1.replies, &d1.attachment)
		if err != nil {
			log.Println("Dz thread 扫描取值失败")
			log.Println(err)
			return
		}

		hydata := hyThread{
			d1.tid,
			d1.fid,
			d1.authorid,
			d1.pid,
			d1.subject,
			d1.dateline,
			d1.lastpost,
			d1.views,
			d1.replies,
			d1.attachment,
		}

		dataArr = append(dataArr, hydata)
	}

	for _, value := range dataArr {
		_, err = stmt.Exec(value.id, value.fid, value.uid, value.pid, value.title, value.atime, value.btime, value.views, value.posts, value.files)
		if err != nil {
			return
		}

		stat++
	}
	defer HybbsDbTx.Rollback()
	HybbsDbTx.Commit()

	if err == nil {
		log.Printf("%s 转换成功, 总共插入 %d 条数据", t.tbname, stat)
	} else {
		log.Printf("%s 转换失败", t.tbname)
	}
	return
}
