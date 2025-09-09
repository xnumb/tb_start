package serv

import (
	"github.com/xnumb/tb"
	tele "gopkg.in/telebot.v4"
	"main/mod"
)

var menus = tb.Menus{
	menuStart,
}

var menuStart = &tb.Menu{
	ID:   "start",
	Desc: "开始",
	Fn: func(c tele.Context) error {
		if !c.Message().Private() {
			return nil
		}
		_, _ = mod.UpdateUser(c.Sender())
		conf, role := checkAuthRole(c)
		if conf == nil {
			return nil
		}
		sp := tb.SendParams{}
		if role == 2 {
			sp.KbBtns = saKbts
			sp.Info = "欢迎超级管理员"
		} else if role == 1 {
			sp.KbBtns = aKbts
			sp.Info = "欢迎管理员"
		} else {
			sp.KbBtns = kbts
			sp.SetMedia(conf.Welcome, conf.WelcomeMedia, "", false)
		}
		return tb.Send(c, sp)
	},
}
