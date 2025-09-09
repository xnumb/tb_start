package app

import "github.com/xnumb/tb/tronGrid"

var Tron *tronGrid.Client

func init() {
	Tron = tronGrid.New(Conf.TronGridKey, Conf.Debug)
}
