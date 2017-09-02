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

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))

	DBConf = readConf()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	//拼接dataSource
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", DBConf.User, DBConf.Pass, DBConf.Url, DBConf.Database)

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	//查询user
	user := User{Id: 2}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		log.Println("没有查询到")
	} else if err == orm.ErrMissPK {
		log.Println("找不到主键")
	} else {
		log.Printf("user:[%d,%s]", user.Id, user.Name)
	}

	//更新name
	user.Name = "jim"
	if num, err := o.Update(&user); err == nil {
		log.Println("update:num=", num)
	} else {
		log.Println("update error:", err)
	}

	//查询默认是通过主键，但是也可以通过指定字段查询，这次通过name来查询
	user2 := User{Name: "tom"}
	log.Println("查询user2:", &user2)
	err = o.Read(&user2, "name")
	if err == orm.ErrNoRows {
		log.Println("user2 没有查询到")
	} else if err == orm.ErrMissPK {
		log.Println("user2 找不到主键")
	} else {
		log.Printf("user2:[%d,%s]", user.Id, user.Name)
	}

	//删除不演示了，见文档
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
