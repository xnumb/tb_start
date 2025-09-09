package serv

import (
	"github.com/xnumb/tb"
	"github.com/xnumb/tb/log"
	tele "gopkg.in/telebot.v4"
	"main/app"
	"main/mod"
)

var Bot *tb.Tbot

func Gen(token string) *tb.Tbot {
	defer func() {
		if r := recover(); r != nil {
			log.Err(nil, "捕获", r)
		}
	}()
	t, err := tb.New(tb.InitParams{
		Token:        token,
		AllowUpdates: tb.AllowUpdatesHigh,
		Proxy:        app.Conf.Proxy,
		BtnExpireMin: app.BtnExpireMin,
		Asker:        mod.Ask{},
		Btns:         btns,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	t.SetMenus(menus, 0)
	t.Client().Handle(tele.OnText, func(c tele.Context) error {
		if c.Message().Private() {
			if kbs.Apply(c) {
				return nil
			}
			if btns.CheckAsker(c, mod.Ask{}) {
				return nil
			}
		}
		return nil
	})
	t.Client().Handle(tele.OnPhoto, func(c tele.Context) error {
		if c.Message().Private() {
			if btns.CheckAsker(c, mod.Ask{}) {
				return nil
			}
		}
		return nil
	})
	t.Client().Handle(tele.OnVideo, func(c tele.Context) error {
		if c.Message().Private() {
			if btns.CheckAsker(c, mod.Ask{}) {
				return nil
			}
		}
		return nil
	})
	Bot = t
	return t
}
