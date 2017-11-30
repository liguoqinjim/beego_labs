package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type Conf struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Url      string `json:"url"`
	Database string `json:"database"`
}

var DBConf *Conf

type Admin struct {
	Id        int
	Name      string
	AdminInfo *AdminInfo `orm:"column(info2_id);rel(one)"` // OneToOne relation
}

type AdminInfo struct {
	Id    int
	Age   int16
	Admin *Admin `orm:"reverse(one)"` //设置一对一反向关系(可选)
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Admin), new(AdminInfo))

	DBConf = readConf()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	//拼接dataSource
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", DBConf.User, DBConf.Pass, DBConf.Url, DBConf.Database)

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	orm.Debug = true

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	adminInfo := new(AdminInfo)
	adminInfo.Age = 33

	admin := new(Admin)
	admin.AdminInfo = adminInfo
	admin.Name = "hello1"

	// 建表
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	//插入
	log.Println(o.Insert(adminInfo))
	log.Println(o.Insert(admin))

	//查询
	admin2 := &Admin{Id: 2}
	o.Read(admin2)
	if admin2.AdminInfo != nil {
		o.Read(admin2.AdminInfo)
		log.Println("admin2=", admin2, admin2.AdminInfo)
	}

	//直接关联查询(这样查询会adminInfo的admin也会赋值，上面那种不会)
	admin3 := &Admin{}
	o.QueryTable("admin").Filter("Id", 3).RelatedSel().One(admin3)
	log.Println("admin3=", admin3)
	log.Println("admin3.AdminInfo=", admin3.AdminInfo.Admin)

	//通过Admin方向查询AdminInfo(通过admin的id,再找到admin的info_id，再找到admin_info)
	var adminInfo4 AdminInfo
	err = o.QueryTable("admin_info").Filter("Admin__Id", 2).One(&adminInfo4)
	if err == nil {
		log.Println("adminInfo4=", adminInfo4)
	}
}

func readConf() *Conf {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("reafFile error:%v", err)
	}

	conf := new(Conf)
	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("unmarshal error:%v", err)
	}

	return conf
}
