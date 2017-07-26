package setting

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"log"
	"os"
)

type Database struct {
	form walk.Form
}

type Hostinfo struct {
	Dbhost,
	Dbuser,
	Dbpwd,
	Dbname,
	Dbport,
	Dbhost2,
	Dbuser2,
	Dbpwd2,
	Dbname2,
	Dbport2 string
}

func (d *Database) Init(owner walk.Form) {
	d.form = owner
	log.Println("database init")
}

func (d *Database) Create() (int, error) {
	log.Println("database Create")

	d.ReadConfig()

	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	animal := new(Hostinfo)

	return Dialog{
		AssignTo:&dlg,
		Title:"配置数据库",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			DataSource:     animal,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout:Grid{Columns:2},
				Children:[]Widget{
					Label{
						Font: Font{Bold: true,Underline: true},
						ColumnSpan: 2,
						Text: "Discuz!7.2数据库信息",
					},
					Label{
						Text:"数据库地址",
					},
					LineEdit{
						Text:Bind("Dbhost"),
					},
					Label{
						Text:"数据库用户名",
					},
					LineEdit{
						Text:Bind("Dbuser"),
					},
					Label{
						Text:"数据库密码",
					},
					LineEdit{
						Text:Bind("Dbpwd"),
					},
					Label{
						Text:"数据库名称",
					},
					LineEdit{
						Text:Bind("Dbname"),
					},
					Label{
						Text:"数据库端口",
					},
					LineEdit{
						Text:Bind("Dbport"),
					},
					Label{
						Font: Font{Bold: true,Underline: true},
						ColumnSpan: 2,
						Text: "Hybbs数据库信息",
					},
					Label{
						Text:"数据库地址",
					},
					LineEdit{
						Text:Bind("Dbhost2"),
					},
					Label{
						Text:"数据库用户名",
					},
					LineEdit{
						Text:Bind("Dbuser2"),
					},
					Label{
						Text:"数据库密码",
					},
					LineEdit{
						Text:Bind("Dbpwd2"),
					},
					Label{
						Text:"数据库名称",
					},
					LineEdit{
						Text:Bind("Dbname2"),
					},
					Label{
						Text:"数据库端口",
					},
					LineEdit{
						Text:Bind("Dbport2"),
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
							log.Printf("%+v", animal)
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
	}.Run(d.form)
}

func (d *Database) ReadConfig() {
	log.Println("ReadConfig 读取文件")
	dbpath := "db.json"
	file, err := os.Open(dbpath) 
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read %d bytes: %q\n", count, data[:count])
}

