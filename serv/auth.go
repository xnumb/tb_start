package client

import (
	"github.com/xnumb/tele"
	"main/mod"
	"strings"
	"tb2/to"
)

func checkAuth(c tele.Context) (*mod.Conf, bool) {
	conf := mod.Conf{}
	if err := conf.Get(); err != nil {
		return nil, false
	}
	if strings.Contains(conf.AdminIds, to.S(c.Sender().ID)) {
		return &conf, true
	}
	return &conf, false
}
func checkSuperAuth(c tele.Context) (*mod.Conf, bool) {
	conf := mod.Conf{}
	if err := conf.Get(); err != nil {
		return nil, false
	}
	if strings.Contains(conf.SuperAdminIds, to.S(c.Sender().ID)) {
		return &conf, true
	}
	return &conf, false
}

// 0:普通用户 1:admin 2:superAdmin
func checkAuthRole(c tele.Context) (*mod.Conf, int) {
	conf := mod.Conf{}
	if err := conf.Get(); err != nil {
		return nil, 0
	}
	if strings.Contains(conf.SuperAdminIds, to.S(c.Sender().ID)) {
		return &conf, 2
	}

	if strings.Contains(conf.AdminIds, to.S(c.Sender().ID)) {
		return &conf, 1
	}
	return &conf, 0
}
