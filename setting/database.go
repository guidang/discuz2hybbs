package setting

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var (
	Data   Dbconf
	animal        = new(Hostinfo)
	dbpath string = "db.json"
)

func (d *Database) Init(owner walk.Form) {
	d.form = owner
	log.Println("database init")
}

func (d *Database) Create() (int, error) {
	log.Println("database Create")

	if err := d.ReadConfig(); err != nil {
		log.Println(err)
	}

	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "配置数据库",
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
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Font:       Font{Family: "微软雅黑", PointSize: 16, Bold: true, Underline: true},
						ColumnSpan: 2,
						Text:       "Discuz!7.2数据库信息",
					},
					Label{
						Text: "数据库地址",
					},
					LineEdit{
						Text: Bind("Dbhost"),
					},
					Label{
						Text: "数据库用户名",
					},
					LineEdit{
						Text: Bind("Dbuser"),
					},
					Label{
						Text: "数据库密码",
					},
					LineEdit{
						Text: Bind("Dbpwd"),
					},
					Label{
						Text: "数据库名称",
					},
					LineEdit{
						Text: Bind("Dbname"),
					},
					Label{
						Text: "数据库端口",
					},
					LineEdit{
						Text: Bind("Dbport"),
					},
					Label{
						Font:       Font{Family: "微软雅黑", PointSize: 16, Bold: true, Underline: true},
						ColumnSpan: 2,
						Text:       "Hybbs数据库信息",
					},
					Label{
						Text: "数据库地址",
					},
					LineEdit{
						Text: Bind("Dbhost2"),
					},
					Label{
						Text: "数据库用户名",
					},
					LineEdit{
						Text: Bind("Dbuser2"),
					},
					Label{
						Text: "数据库密码",
					},
					LineEdit{
						Text: Bind("Dbpwd2"),
					},
					Label{
						Text: "数据库名称",
					},
					LineEdit{
						Text: Bind("Dbname2"),
					},
					Label{
						Text: "数据库端口",
					},
					LineEdit{
						Text: Bind("Dbport2"),
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
							d.WriteConfig()

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

func (d *Database) ReadConfig() (err error) {
	log.Println("ReadConfig 读取文件")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		log.Println("数据库配置文件不存在")
		return err
	}

	bytes, err := ioutil.ReadFile(dbpath)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("读取的数据:\n%s\n", bytes)

	//dataStr := fmt.Sprintf("%s", data)
	//log.Println(dataStr)
	if err := json.Unmarshal(bytes, &Data); err != nil {
		log.Println("Json转Struct出错")
		log.Println(err)
		return err
	}

	log.Println(Data)

	animal.Dbhost = Data.Discuz.Dbhost
	animal.Dbuser = Data.Discuz.Dbuser
	animal.Dbpwd = Data.Discuz.Dbpwd
	animal.Dbname = Data.Discuz.Dbname
	animal.Dbport = Data.Discuz.Dbport

	animal.Dbhost2 = Data.Hybbs.Dbhost
	animal.Dbuser2 = Data.Hybbs.Dbuser
	animal.Dbpwd2 = Data.Hybbs.Dbpwd
	animal.Dbname2 = Data.Hybbs.Dbname
	animal.Dbport2 = Data.Hybbs.Dbport

	return err
}

func (d *Database) WriteConfig() (err error) {

	Data.Discuz.Dbhost = animal.Dbhost
	Data.Discuz.Dbuser = animal.Dbuser
	Data.Discuz.Dbpwd = animal.Dbpwd
	Data.Discuz.Dbname = animal.Dbname
	Data.Discuz.Dbport = animal.Dbport

	Data.Hybbs.Dbhost = animal.Dbhost2
	Data.Hybbs.Dbuser = animal.Dbuser2
	Data.Hybbs.Dbpwd = animal.Dbpwd2
	Data.Hybbs.Dbname = animal.Dbname2
	Data.Hybbs.Dbport = animal.Dbport2

	dataByte, err := json.Marshal(Data)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(dbpath, dataByte, 0755)

	if err != nil {
		log.Println(err)
	}

	return err
}

type Dbinfo struct {
	Dbhost string `json:"dbhost"`
	Dbuser string `json:"dbuser"`
	Dbpwd  string `json:"dbpwd"`
	Dbname string `json:"dbname"`
	Dbport string `json:"dbport"`
}

type Dbconf struct {
	Discuz,
	Hybbs Dbinfo
}
