package setting

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Config struct {
	Animal Info
	Form   walk.Form
}

type Info struct {
	Adminid string
}

func (c *Config) Create() (code int, err error) {
	log.Println("config Create")

	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	code, err = Dialog{
		AssignTo:      &dlg,
		Title:         "基本配置",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			DataSource:     &c.Animal,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 100},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Font:       Font{Family: "微软雅黑", PointSize: 16, Bold: true, Underline: true},
						ColumnSpan: 2,
						Text:       "基本配置",
					},
					Label{
						Text: "管理员ID",
					},
					LineEdit{
						Text: Bind("Adminid"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "确定",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
							//log.Printf("%+v", c.Animal)
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "取消",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(c.Form)

	return
}
