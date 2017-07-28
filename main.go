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
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction *walk.Action

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Discuz转Hybbs",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
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
						Text: "配置数据库",
						OnClicked: func() {
							db := new(setting.Database)
							db.Init(mw)
							if cmd, err := db.Create(); err != nil {
								log.Print(err)
							} else if cmd == walk.DlgCmdOK {

							}
						},
					},
					PushButton{
						Text: "配置管理员",
						OnClicked: func() {
							log.Println("点击配置管理员")
						},
					},
				},
			},
			PushButton{
				Text: "开始转换",
				OnClicked: func() {
					log.Println("点击开始转换")
					convert := new(model.Convert)
					convert.Init(mw)
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{300, 200},
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
	var msg string = "作者: Skiychan\r\n邮箱: dev@skiy.net\r\n网站: https://www.skiy.net\r\n版本: 0.0.1"
	walk.MsgBox(mw, "关于", msg, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
