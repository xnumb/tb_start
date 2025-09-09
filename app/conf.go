package app

import (
	"flag"
	"github.com/xnumb/tb/log"
	"github.com/xnumb/tb/utils"
)

type config struct {
	Debug bool   `yaml:"debug"`
	Proxy string `yaml:"proxy"`
	Token string `yaml:"token"`
	DB    struct {
		User string `yaml:"user"`
		Pwd  string `yaml:"pwd"`
		Name string `yaml:"name"`
	} `yaml:"db"`
}

var Conf *config

const (
	BtnExpireMin = 0  // 0代表禁用过期检测
	AskExpireMin = 10 // 0代表不过期
)

func init() {
	confPath := flag.String("conf", "conf.yaml", "path to config file")
	flag.Parse()
	if err := utils.ParseYaml(*confPath, &Conf); err != nil {
		log.Fatal(err)
	}
}
