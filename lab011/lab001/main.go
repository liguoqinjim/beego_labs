package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type Conf struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Url      string `json:"url"`
	Database string `json:"database"`
}

var DBConf *Conf

func init() {
	DBConf = readConf()
}

func main() {
	log.Println(DBConf)
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
