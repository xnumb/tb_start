package mod

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Tid       int64  `gorm:"index"`
	Name      string `gorm:"type:varchar(64)"`
	Token     enum.Token
	AgentId   uint `gorm:"index"`
	Price     tb.Amount
	Welcome   string `gorm:"type:text"`
	Index     string `gorm:"type:varchar(64);default:'1';index;unique;not null"`
	IsDone    bool
	Expire    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Agent     *Agent    `gorm:"-"`
}

type Users []*User

type UserPage struct {
	Total int64
	Count int
	List  Users
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

func (r *User) Del(id uint) error {
	if err := r.Get(id); err != nil {
		return err
	}
	return db.Delete(&r).Error
}

func (r *User) Save() error {
	if r.ID == 0 {
		return ErrModSaveNoPrimaryKey
	}
	return db.Save(&r).Error
}
func (rs *Users) Get() error {
	return db.Find(&rs).Error
}

func (p *UserPage) Get(page, size int) error {
	if err := db.Model(&User{}).Count(&p.Total).Error; err != nil {
		return err
	}
	p.Count = CalcPageCount(p.Total, size)
	start := (page - 1) * size
	if err := db.Order("created_at desc").Limit(size).Offset(start).Find(&p.List).Error; err != nil {
		return err
	}
	return nil
}
