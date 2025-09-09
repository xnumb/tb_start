package mod

import (
	"github.com/xnumb/tele"
	"strings"
	"time"
)

type User struct {
	ID        uint
	Tid       int64  `gorm:"index"`
	Name      string `gorm:"type:varchar(64)"`
	Username  string `gorm:"type:varchar(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *User) GetByTid(tid int64) error {
	return db.First(&r, "tid = ?", tid).Error
}

func (r *User) Get(id uint) error {
	return db.First(&r, id).Error
}

func (r *User) Add() error {
	return db.Create(&r).Error
}

func (r *User) Save() error {
	if r.ID == 0 {
		return ErrModSaveNoPrimaryKey
	}
	return db.Save(&r).Error
}

func UpdateUser(u *tele.User) (*User, error) {
	if u == nil {
		return nil, ErrNoUserInstance
	}
	name := strings.TrimSpace(u.FirstName + " " + u.LastName)
	r := User{}
	if err := r.GetByTid(u.ID); err != nil {
		r2 := User{
			Tid:      u.ID,
			Name:     name,
			Username: u.Username,
		}
		if err = r2.Add(); err != nil {
			return nil, err
		}
		return &r2, nil
	}
	r.Name = name
	r.Username = u.Username
	if err := r.Save(); err != nil {
		return nil, err
	}
	return &r, r.Save()
}
