package model

import (
	"github.com/lxn/walk"
	//. "github.com/lxn/walk/declarative"
)

import (
	"log"
	"../setting"
	"database/sql"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbinfo setting.Dbinfo
	dbconf setting.Dbconf
	DiscuzDb *sql.DB
	HybbsDb *sql.DB
)

type Convert struct {
	form walk.Form
}

func (c *Convert) Init(owner walk.Form) {
	c.form = owner
	log.Println("Convert init")
}

func (c *Convert) ReadConfig() {
	log.Println("ReadConfig 读取文件")

	dbpath := "db.json"
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		log.Println("数据库配置文件不存在")
		return
	}

	bytes, err := ioutil.ReadFile(dbpath)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("读取的数据:\n%s\n", bytes)

	//dataStr := fmt.Sprintf("%s", data)
	//log.Println(dataStr)
	if err := json.Unmarshal(bytes, &dbconf); err != nil {
		log.Println("Json转Struct出错")
		log.Println(err)
		return
	}
}

func (c *Convert) CheckConnect(flag int) (db *sql.DB, err error) {
	log.Println("CheckConnect 检测数据库连接")

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
	log.Println("dbStr: " + dbStr)

	db, err1 := sql.Open("mysql", dbStr)
	log.Println(err1)
	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		err = errors.New("数据库连接失败")
	}
	return
}

func (c *Convert) ToHybbs() (err error){
	c.ReadConfig()

	DiscuzDb, err = c.CheckConnect(1)
	if err != nil {
		log.Println(err)
		log.Println("Discuz数据库连接失败")
	}

	HybbsDb, err = c.CheckConnect(2)
	if err != nil {
		log.Println(err)
		log.Println("Hybbs数据库连接失败")
	}

	return
}

func (c *Convert) Test() (err error){
	log.Println("Test")
	rows, err := DiscuzDb.Query("SELECT * FROM pre_common_member")
	if err != nil {
		log.Println("test")
		log.Println(err)
		return
	}
	var t int
	for rows.Next() {
		t += 1
		log.Println(t)
	}
	return err
}
