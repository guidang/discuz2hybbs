package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/skiy/discuz2hybbs/model"
	"github.com/skiy/discuz2hybbs/setting"
)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var isSpecialMode = walk.NewMutableCondition()

var cf setting.Config

type versionData struct {
	Code int
	Msg  string
}

const (
	version = "0.1.0"
)

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var showCheckVersionBoxAction, showAboutBoxAction, showAboutAuthorAction *walk.Action
	var te *walk.TextEdit

	cf.Form = mw

	if err := (MainWindow{
		//Icon:     "dh.ico",
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
						AssignTo:    &showCheckVersionBoxAction,
						Text:        "检测更新",
						OnTriggered: mw.showCheckVersion_Triggered,
					},
					Action{
						AssignTo:    &showAboutAuthorAction,
						Text:        "关于作者",
						OnTriggered: mw.showAboutAuthorAction_Triggered,
					},
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "关于软件",
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
							db := setting.Database{Form: mw}
							if cmd, err := db.Create(); err != nil {
								log.Print(err)
							} else if cmd == walk.DlgCmdOK {

							}
						},
					},
					PushButton{
						Text: "基本配置",
						OnClicked: func() {
							var cmd int
							var err error
							cmd, err = cf.Create()

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
					//log.Println("点击开始转换")
					convert := setting.Convert{
						cf.Animal,
						mw,
						te,
					}
					//te.SetText("正在转换...")
					//log.Println(convert)
					convert.Create()
				},
			},
			TextEdit{
				AssignTo: &te,
				ReadOnly: true,
				MinSize:  Size{300, 240},
				VScroll:  true,
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showCheckVersionBoxAction},
			ActionRef{&showAboutAuthorAction},
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

func (mw *MyMainWindow) showCheckVersion_Triggered() {
	url := fmt.Sprintf("https://www.skiy.net/soft/checkVersion.php?version=%s&name=%s", version, "dz2hy")
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var dataStr versionData
	var message string
	var infoStyle walk.MsgBoxStyle

	json.Unmarshal(body, &dataStr)

	if dataStr.Code == 0 {
		message = `已经是最新版本`
		infoStyle = walk.MsgBoxIconInformation
	} else if dataStr.Code == -1 {
		message = dataStr.Msg
		infoStyle = walk.MsgBoxIconError
	} else {
		message = dataStr.Msg
		infoStyle = walk.MsgBoxIconWarning
	}

	walk.MsgBox(mw, "检测更新", message, infoStyle)
}

func (mw *MyMainWindow) showAboutAuthorAction_Triggered() {
	var msg string = `
作者: Skiychan
Q Q:  86999070
邮箱: dev@skiy.net
网站: https://www.skiy.net`
	walk.MsgBox(mw, "关于作者", msg, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	var msg string = `
Version: 0.1.0

Project: https://github.com/skiy/discuz2hybbs
`
	walk.MsgBox(mw, "关于软件", msg, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
