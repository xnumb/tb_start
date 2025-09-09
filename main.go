package main

import (
	"github.com/xnumb/tb/log"
	"main/app"
	"main/mod"
	"main/serv"
)

//import "github.com/robfig/cron/v3"

func main() {
	run()
}

func run() {
	mod.Reset()

	//c := cron.New()
	//_, _ = c.AddFunc("@every 5s", func() {
	//
	//})
	//c.Start()

	log.Info("Start client bot.")
	b := serv.Gen(app.Conf.Token)
	if b != nil {
		if err := (&mod.Conf{}).SaveBotInfo(b.Client().Me); err != nil {
			log.Err(err)
		}
		b.Start()
	}
}
