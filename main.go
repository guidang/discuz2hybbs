// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"./model"
	"./setting"
	"fmt"
	"time"
)

var isSpecialMode = walk.NewMutableCondition()

var info setting.Info

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var showAboutBoxAction *walk.Action

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Discuz转Hybbs",
		MenuItems: []MenuItem{
			Menu{
				Text: "&菜单",
				Items: []MenuItem{
					Separator{},
					Action{
						Text:        "退出",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&帮助",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "关于",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "数据库配置",
						OnClicked: func() {
							db := new(setting.Database)
							db.Form = mw
							if cmd, err := db.Create(); err != nil {
								log.Print(err)
							} else if cmd == walk.DlgCmdOK {

							}
						},
					},
					PushButton{
						Text: "基本配置",
						OnClicked: func() {
							cf := new(setting.Config)
							cf.Form = mw
							var cmd int
							var err error
							cmd, err, info = cf.Create()

							if err != nil {
								log.Print(err)
							} else if cmd == walk.DlgCmdOK {
							}
						},
					},
				},
			},
			PushButton{
				Text: "开始转换",
				OnClicked: func() {
					log.Println("点击开始转换")
					t1 := time.Now()
					convert := model.Convert{
						info,
						mw,
					}

					err := convert.ToHybbs()
					if err == nil {
						t2 := time.Now()
						d := t2.Sub(t1)
						fmt.Printf("\r\n已经成功将 Discuz 转换成 Hybbs, 总共耗时: %s\r\n", d)
					}
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{300, 120},
		Layout:  VBox{},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	var msg string = `
作者: Skiychan
Q Q:  1005043848
邮箱: dev@skiy.net
网站: https://www.skiy.net
版本: 0.0.1`
	walk.MsgBox(mw, "关于", msg, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
