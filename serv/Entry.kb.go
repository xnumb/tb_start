package client

import (
	"github.com/xnumb/tele"
	tb "tb2"
	"tb2/emj"
)

var kbs = tb.Kbs{
	kbConf,
	kbAdmins,
}

var kbts = tb.Kbts{}

var aKbts = tb.Kbts{
	{_kbConf},
}

var saKbts = tb.Kbts{
	{_kbConf, _kbAdmins},
}

const _kbConf = emj.Gear + " 系统配置"

var kbConf = &tb.Kb{
	Text: _kbConf,
	Fn: func(c tele.Context) error {
		return sendConf(c, false)
	},
}

const _kbAdmins = emj.User + " 管理员管理"

var kbAdmins = &tb.Kb{
	Text: _kbAdmins,
	Fn: func(c tele.Context) error {
		return sendAdmins(c, false)
	},
}
