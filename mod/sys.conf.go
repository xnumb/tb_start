package mod

import (
	"tb2/log"
	"time"
)

// All table struct
var autoMigrateDst = []any{
	&Ask{},
	&Conf{},
	&User{},
}

// Invoked when db reset
func initDb() {
	config := Conf{}
	err := config.AddIfNotExist()
	if err != nil {
		log.Err(err, "初始化Conf失败")
	}
}

type Conf struct {
	ID            uint
	BotTid        int64
	BotName       string `gorm:"type:varchar(64)"`
	BotUsername   string `gorm:"type:varchar(64)"`
	SuperAdminIds string `gorm:"type:varchar(256)"`
	AdminIds      string `gorm:"type:varchar(256)"`
	Welcome       string `gorm:"type:text"`
	WelcomeMedia  string `gorm:"type:varchar(128)"`
	UpdatedAt     time.Time
}
