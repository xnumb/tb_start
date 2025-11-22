package serv

import (
	"main/mod"
	"strings"

	"github.com/xnumb/tb/to"
	tele "gopkg.in/telebot.v4"
)

func checkAuth(c tele.Context) (*mod.Conf, bool) {
	conf := mod.Conf{}
	if err := conf.Get(); err != nil {
		return nil, false
	}
	senderId := to.S(c.Sender().ID)
	if strings.Contains(conf.AdminIds, senderId) {
		return &conf, true
	}
	if strings.Contains(conf.SuperAdminIds, senderId) {
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

func getAllAdmins() []int64 {
	conf := mod.Conf{}
	if err := conf.Get(); err != nil {
		return nil
	}
	admins := strings.Split(conf.AdminIds, ",")
	superAdmins := strings.Split(conf.SuperAdminIds, ",")
	allAdmins := append(admins, superAdmins...)

	// 去重
	seen := make(map[string]struct{})
	unique := []int64{}
	for _, v := range allAdmins {
		if v == "" {
			continue
		}
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			val, ok := to.I64(v)
			if !ok {
				continue
			}
			unique = append(unique, val)
		}
	}
	return unique
}
