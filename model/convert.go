package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lxn/walk"
	"github.com/skiy/discuz2hybbs/setting"
)

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	dbinfo     setting.Dbinfo
	dbconf     setting.Dbconf
	DiscuzDb   *sql.DB
	HybbsDb    *sql.DB
	DiscuzDbTx *sql.Tx
	HybbsDbTx  *sql.Tx
)

type Convert struct {
	Info setting.Info
	Form walk.Form
	Te   *walk.TextEdit
}

type Msg struct {
	code int
	log  string
}

var (
	Te *walk.TextEdit
)

func (c *Convert) Create() {
	t1 := time.Now()

	Te = c.Te
	err := c.ToHybbs()
	t2 := time.Now()
	d := t2.Sub(t1)

	var msg string
	if err != nil {
		msg = "\r\n\r\n=================================\r\nDiscuz 转换成 Hybbs 失败，请自行检查数据库配置\r\n【开发者】\r\nQQ: 1005043848\r\nEmail: dev@skiy.net\r\n意见反馈: https://github.com/skiy/DiscuzToHybbs\r\n\r\n"
	} else {
		msg = fmt.Sprintf("\r\n已经成功将 Discuz 转换成 Hybbs，总共耗时: %s\r\n", d)
	}

	msg = Te.Text() + msg
	Te.SetText(msg)
}

func (c *Convert) ReadConfig() (err error) {
	SetConvertLog("读取数据库配置信息...", 0)

	dbpath := "db.json"
	if _, err = os.Stat(dbpath); os.IsNotExist(err) {
		SetConvertLog("数据库配置文件不存在", -1)
		log.Println(err)
		return
	}

	bytes, err := ioutil.ReadFile(dbpath)
	if err != nil {
		log.Println(err)
		return
	}

	//fmt.Printf("读取的数据:\n%s\n", bytes)

	//dataStr := fmt.Sprintf("%s", data)
	//log.Println(dataStr)
	if err = json.Unmarshal(bytes, &dbconf); err != nil {
		SetConvertLog("Json转Struct出错", -1)
		log.Println(err)
		return
	}

	return nil
}

func (c *Convert) CheckConnect(flag int) (db *sql.DB, err error) {
	//log.Println("CheckConnect 检测数据库连接")

	//log.Println(dbconf)

	if flag == 1 {
		dbinfo = dbconf.Discuz
	} else {
		dbinfo = dbconf.Hybbs
	}

	if dbinfo.Dbhost == "" {
		err = errors.New("数据库地址不能为空")
		return
	}

	if dbinfo.Dbuser == "" {
		err = errors.New("数据库用户名不能为空")
		return
	}

	if dbinfo.Dbname == "" {
		err = errors.New("数据库名称不能为空")
		return
	}

	hostStr := dbinfo.Dbhost
	if dbinfo.Dbport != "" {
		hostStr = "tcp(" + hostStr + ":" + dbinfo.Dbport + ")"
	}

	dbStr := fmt.Sprintf("%s:%s@%s/%s?%s",
		dbinfo.Dbuser,
		dbinfo.Dbpwd,
		hostStr,
		dbinfo.Dbname,
		"utf8",
	)
	//log.Println("dbStr: " + dbStr)

	db, err = sql.Open("mysql", dbStr)
	//log.Println(err)
	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		err = errors.New("数据库连接失败")
	}
	return
}

func (c *Convert) Tx() {
	var err error
	DiscuzDbTx, err = DiscuzDb.Begin()
	if err != nil {
		log.Println("Discuz 事务 Begin 失败")
		log.Println(err)
		return
	}
	HybbsDbTx, err = HybbsDb.Begin()
	if err != nil {
		log.Println("Hybbs 事务 Begin 失败")
		log.Println(err)
		return
	}
	return
}

func (c *Convert) ToHybbs() (err error) {
	c.ReadConfig()
	DiscuzDb, err = c.CheckConnect(1)
	if err != nil {
		SetConvertLog(fmt.Sprintf("Discuz数据库连接失败，具体原因：\r\n%s", err.Error()), -1)
		log.Println(err)
		return
	}

	HybbsDb, err = c.CheckConnect(2)
	if err != nil {
		SetConvertLog(fmt.Sprintf("Hybbs数据库连接失败，具体原因：\r\n%s", err.Error()), -1)
		log.Println(err)
		return
	}

	//版块转换
	f := new(Forum)
	c.Tx()
	err = f.Init()
	if err != nil {
		log.Println(err)
		return
	}

	//主题转换
	t := new(Thread)
	c.Tx()
	err = t.Init()
	if err != nil {
		log.Println(err)
		return
	}

	//帖子转换
	p := new(Post)
	c.Tx()
	err = p.Init()
	if err != nil {
		log.Println(err)
		return
	}

	//用户转换
	u := new(User)
	u.adminid = c.Info.Adminid
	c.Tx()
	err = u.Init()
	if err != nil {
		log.Println(err)
		return
	}

	SetConvertLog("", 2)
	return nil
}
