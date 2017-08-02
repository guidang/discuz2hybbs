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
	SetConvertLog("正在转换 "+t.tbname+" ...", 0)

	err = Truncate(t.tbname)
	if err != nil {
		return
	}

	dzSqlStr := "SELECT t.tid, t.fid, t.authorid, p.pid, t.subject, t.dateline, t.lastpost, t.views, t.replies, t.attachment FROM `cdb_threads` t LEFT JOIN cdb_posts p ON p.tid = t.tid WHERE p.first = 1"
	hySqlStr := fmt.Sprintf("INSERT INTO %s (id, fid, uid, pid, title, atime, btime, views, posts, files, summary, img) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, '', '')", t.tbname)

	data, err := DiscuzDbTx.Query(dzSqlStr)
	if err != nil {
		SetConvertLog("Dz thread 查询失败: "+dzSqlStr, -1)
		log.Println(err)
		return
	}

	stmt, err := HybbsDbTx.Prepare(hySqlStr)
	if err != nil {
		SetConvertLog("Hy thread 预加载失败: "+dzSqlStr, -1)
		log.Println(err)
		return
	}

	var stat int
	var dataArr []hyThread

	for data.Next() {
		d1 := new(dzThread)
		err = data.Scan(&d1.tid, &d1.fid, &d1.authorid, &d1.pid, &d1.subject, &d1.dateline, &d1.lastpost, &d1.views, &d1.replies, &d1.attachment)
		if err != nil {
			SetConvertLog("Dz thread 扫描取值失败", -1)
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
		msg := fmt.Sprintf("%s 转换成功, 总共插入 %d 条数据", t.tbname, stat)
		SetConvertLog(msg, 0)
	} else {
		msg := fmt.Sprintf("%s 转换失败")
		SetConvertLog(msg, -1)
		log.Println(err)
	}
	return
}
