package serv

import (
	"github.com/xnumb/tb"
	"github.com/xnumb/tb/emj"
	"github.com/xnumb/tb/to"
	tele "gopkg.in/telebot.v4"
	"strings"
)

var _btnConf = &tb.Btn{
	ID:   "conf",
	Args: []string{"name"},
}

var btnConf = _btnConf.LinkAsk(tb.AskAction{
	Q: func(c tele.Context, args tb.Args) string {
		name := args.Get("name")
		switch name {
		case "welcome":
			return "请输入机器人欢迎语(可以包含一个图片或一个视频)"
		default:
			return "未定义"
		}
	},
	Fn: func(c tele.Context, val string, args tb.Args) (bool, error) {
		conf, ok := checkAuth(c)
		if !ok {
			return true, nil
		}
		name := args.Get("name")

		if name == "welcome" {
			text, mediaId := tb.ReceiveMedia(c)
			conf.Welcome = text
			conf.WelcomeMedia = mediaId
		}
		if err := conf.Save(); err != nil {
			return true, tb.SendErr(c, err)
		}
		return true, sendConf(c, false)
	},
})

var _btnConfPreview = &tb.Btn{
	ID:   "confPreview",
	Text: emj.Pic + " 预览",
	Args: []string{"name"},
}

var btnConfPreview = _btnConfPreview.Link(tb.Action{
	Fn: func(c tele.Context, args tb.Args) error {
		r, ok := checkAuth(c)
		if !ok {
			return nil
		}
		name := args.Get("name")
		sp := tb.SendParams{}
		if name == "welcome" {
			sp.SetMedia(r.Welcome, r.WelcomeMedia, "", true)
		}
		return tb.Send(c, sp)
	},
})

var _btnAdminDel = &tb.Btn{
	ID:   "adminDel",
	Text: emj.X + " 删除",
	Args: []string{"id"},
}

var btnAdminDel = _btnAdminDel.Link(tb.Action{
	Fn: func(c tele.Context, args tb.Args) error {
		conf, ok := checkSuperAuth(c)
		if !ok {
			return nil
		}
		id := args.Get("id")
		admins := strings.Split(conf.AdminIds, ",")
		for i, r := range admins {
			if r == id {
				admins = append(admins[:i], admins[i+1:]...)
			}
		}
		conf.AdminIds = strings.Join(admins, ",")
		if err := conf.Save(); err != nil {
			return tb.SendErr(c, err)
		}
		return sendAdmins(c, true)
	},
})

var _btnAdminAdd = &tb.Btn{
	ID:   "adminAdd",
	Text: emj.Plus + " 新增管理员",
}

var btnAdminAdd = _btnAdminAdd.LinkAsk(tb.AskAction{
	Q: func(c tele.Context, args tb.Args) string {
		return "请输入管理员的飞机ID"
	},
	Fn: func(c tele.Context, val string, args tb.Args) (bool, error) {
		conf, ok := checkSuperAuth(c)
		if !ok {
			return true, nil
		}
		_, ok = to.I64(val)
		if !ok {
			return false, c.Reply("请输入正确的飞机ID")
		}
		admins := strings.Split(conf.AdminIds, ",")
		for _, r := range admins {
			if r == val {
				return true, c.Reply("该用户已是管理员")
			}
		}
		admins = append(admins, val)
		conf.AdminIds = strings.Join(admins, ",")
		if err := conf.Save(); err != nil {
			return true, tb.SendErr(c, err)
		}
		return true, sendAdmins(c, false)

	},
})
