package main

import (
	"bufio"
	"fmt"
	"github.com/skiy/discuz2hybbs/model"
	"github.com/skiy/golib"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	fmt.Println(`
:::
::: 本程序开源地址: https://github.com/skiy/discuz2hybbs
::: 作者: Skiychan <dev@skiy.net> https://www.skiy.net
:::
::: 请选择主菜单:::
:::
::: 1. Discuz!7.2 转换到 HYBBS2
:::
::: 执行过程中按"Q", 再按"回车键"退出本程序...
:::
::: Version:1.0.0    Updated:2018-04-09
`)

	inputVal := lib.Input(buf)

	if inputVal == "1" {
		app := model.App{}
		app.Init()
	}
}
