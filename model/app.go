package model

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/skiy/golib"
	"log"
	"os"
	"strings"
)

type App struct {
}

type dbstr struct {
	lib.Database
	DBPre string
	Auto  bool
}

var (
	DiscuzDb *sql.DB
	HybbsDb  *sql.DB
)

func (this *App) Init() {
	fmt.Println("\r\n===您选择了: 1. Discuz!7.2 转换到 HYBBS2\r\n")

	db3str := dbstr{}
	fmt.Println("正在配置 Discuz 数据库")
	db3str.Setting()

	buf := bufio.NewReader(os.Stdin)
	fmt.Println("请配置数据库表前缀:(默认为 cdb_)")
	s := lib.Input(buf)
	if s == "" {
		s = "cdb_"
	}
	db3str.DBPre = s
	fmt.Println("数据库表前缀为: " + s)

	var err error
	DiscuzDb, err = db3str.Connect()
	if err != nil {
		fmt.Println(err)
		log.Fatalln("\r\nDiscuz 数据库配置错误")
	}

	err = DiscuzDb.Ping()
	if err != nil {
		log.Fatalln("\r\nDiscuz: " + err.Error())
	}

	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	db4str := dbstr{}
	fmt.Println("正在配置 HYBBS 数据库")
	db4str.Setting()

	buf = bufio.NewReader(os.Stdin)
	fmt.Println("请配置数据库表前缀:(默认为 hy_)")
	s = lib.Input(buf)
	if s == "" {
		s = "hy_"
	}
	db4str.DBPre = s
	fmt.Println("数据库表前缀为: " + s)

	HybbsDb, err = db4str.Connect()
	if err != nil {
		log.Fatalln("\r\nHYBBS 数据库配置错误")
	}

	err = HybbsDb.Ping()
	if err != nil {
		log.Fatalln("\r\nHYBBS: " + err.Error())
	}

	if db3str.DBHost == db4str.DBHost && db3str.DBName == db4str.DBName {
		log.Fatalln("\r\n不能在同一个数据库里升级，否则数据会被清空！请将新论坛安装到其他数据库。")
	}

	DiscuzDb.SetMaxIdleConns(0)
	HybbsDb.SetMaxIdleConns(0)

	buf = bufio.NewReader(os.Stdin)
	fmt.Println("全自动更新所有表(Y/N): (默认为 Y)")
	s = lib.Input(buf)
	if !strings.EqualFold(s, "N") {
		db4str.Auto = true
	}

	tables := [...]string{
		//"forum",
		//"thread",
		//"post",
		"user",
	}

	for _, table := range tables {
		fmt.Println("正在转换表: " + table)

		switch table {
		case "forum":
			do := Forum{}
			do.Init()
			break

		case "thread":
			do := Thread{}
			do.Init()
			break

		case "post":
			do := Post{}
			do.Init()
			break

		case "user":
			do := User{}
			do.Init()
			break
		}
	}

	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(`
:::
::: https://github.com/skiy/discuz2hybbs
::: 本程序开源地址: https://github.com/skiy/xiuno-tools
::: 作者: Skiychan <dev@skiy.net> https://www.skiy.net
:::
`)
}
