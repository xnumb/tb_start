package serv

import (
	"github.com/xnumb/tb"
	"github.com/xnumb/tb/emj"
	tele "gopkg.in/telebot.v4"
	"strings"
)

func sendConf(c tele.Context, isEdit bool) error {
	_, ok := checkAuth(c)
	if !ok {
		return nil
	}
	return tb.Send(c, tb.SendParams{
		IsEdit: isEdit,
		Info:   tb.Info("系统配置"),
		Rows: []tele.Row{
			{
				_btnConf.T(emj.Celebrate+" 设置欢迎语", "welcome"),
				_btnConfPreview.G("welcome"),
			},
		},
	})
}

func sendAdmins(c tele.Context, isEdit bool) error {
	conf, ok := checkSuperAuth(c)
	if !ok {
		return nil
	}
	admins := strings.Split(conf.AdminIds, ",")

	var rows []tele.Row
	for _, r := range admins {
		if strings.Contains(conf.SuperAdminIds, r) {
			rows = append(rows, tele.Row{
				tele.Btn{
					Text: r,
					Data: "-",
				},
			})
		} else {
			rows = append(rows, tele.Row{
				tele.Btn{
					Text: r,
					Data: "-",
				},
				_btnAdminDel.G(r),
			})
		}
	}

	return tb.Send(c, tb.SendParams{
		IsEdit: isEdit,
		Info:   tb.Info("管理员管理"),
		Rows:   rows,
		FootRows: []tele.Row{
			{
				_btnAdminAdd.G(),
			},
		},
	})
}
