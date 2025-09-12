package mod

import (
	"errors"
	"fmt"
	"main/app"
	"strings"
	"time"

	"github.com/xnumb/tb/log"
	"github.com/xnumb/tb/utils"
	tele "gopkg.in/telebot.v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB Base
var db *gorm.DB

var (
	ErrModSaveNoPrimaryKey = errors.New("不存在该主键")
	ErrNoUserInstance      = errors.New("不存在User实例")
)

func Init() {
	dbConf := app.Conf.DB
	s := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Asia%sShanghai", dbConf.User, dbConf.Pwd, dbConf.Name, "%2F")
	d, err := gorm.Open(mysql.Open(s), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			//return time.Now().Local()
			return utils.GetNow()
		},
	})
	if err != nil {
		log.Fatal(err, "初始化数据库失败")
	}
	db = d
}

func Close() {
	d, err := db.DB()
	if err != nil {
		log.Err(err, "数据库关闭时未找到db")
		return
	}
	err = d.Close()
	if err != nil {
		log.Err(err, "数据库关闭错误")
	}
}

func SQL(res any, sql string, params ...any) {
	db.Raw(sql, params...).Scan(res)
}

func Reset() {
	Init()
	err := db.AutoMigrate(
		autoMigrateDst...,
	)
	if err != nil {
		log.Fatal(err, "Reset数据库错误")
	}
	initDb()
}

func CalcPageCount(total int64, size int) int {
	s := int64(size)
	c := int(total / s)
	if total%s != 0 {
		c++
	}
	return c
}

// Ask funcs

type Ask struct {
	SenderId  int64  `gorm:"primaryKey"`
	Cmd       string `gorm:"type:varchar(256)"`
	MessageId int
	UpdatedAt time.Time
}

func (r Ask) Set(senderId int64, cmd string, messageId int) error {
	err := db.First(&r, "sender_id = ?", senderId).Error
	r.Cmd = cmd
	r.MessageId = messageId
	if err == nil {
		return db.Save(&r).Error
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.SenderId = senderId
			return db.Create(&r).Error
		}
		return err
	}
}

func (r Ask) Get(senderId int64) string {
	// 先删除超过一定时间的
	if app.AskExpireMin > 0 {
		t := utils.GetNow().Add(-time.Minute * app.AskExpireMin)
		db.Where("updated_at < ?", t).Delete(&Ask{})
	}
	err := db.First(&r, "sender_id = ?", senderId).Error
	if err != nil {
		return ""
	}
	return r.Cmd
}

func (r Ask) Done(senderId int64) (int, error) {
	err := db.First(&r, senderId).Error
	if err != nil {
		return 0, err
	} else {
		msgId := r.MessageId
		err := db.Delete(&r).Error
		if err != nil {
			return 0, err
		}
		return msgId, nil
	}
}

// Conf funcs

func (r *Conf) AddIfNotExist() error {
	err := db.First(r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(r).Error
	} else {
		return err
	}
}

func (r *Conf) Save() error {
	if r.ID == 0 {
		return ErrModSaveNoPrimaryKey
	}
	return db.Save(r).Error
}

func (r *Conf) Get() error {
	return db.First(r).Error
}

func (r *Conf) SaveBotInfo(b *tele.User) error {
	if b == nil {
		return nil
	}
	if err := r.Get(); err != nil {
		return err
	}
	r.BotName = strings.TrimSpace(b.FirstName + " " + b.LastName)
	r.BotUsername = b.Username
	r.BotTid = b.ID
	return r.Save()
}
