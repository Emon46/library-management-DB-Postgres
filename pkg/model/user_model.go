package model

import (
	"time"
)

type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:_`
	PhoneNo  string `json:"phone_no"`
	UserType string `json:"user_type"`
}

type UserDb struct {
	ID       int    `xorm:"pk autoincr id"`
	Name     string `xorm:"name"`
	Mail     string `xorm:"mail"`
	Password string `xorm:"password"`
	PhoneNo  string `xorm:"phone_no"`
	UserType string `xorm:"user_type"`

	CreatedAt time.Time `xorm:"created" `
	UpdatedAt time.Time `xorm:"updated"`
}

func APIFormat(u UserDb) *UserModel {
	return &UserModel{
		ID:       u.ID,
		Name:     u.Name,
		Mail:     u.Mail,
		PhoneNo:  u.PhoneNo,
		UserType: u.UserType,
		Password: "private",
	}
}
func UserDbFormat(user UserModel) UserDb {
	return UserDb{
		ID:       user.ID,
		Name:     user.Name,
		Mail:     user.Mail,
		Password: user.Password,
		PhoneNo:  user.PhoneNo,
		UserType: user.UserType,
	}

}

type UsersDb struct {
	Users []UserDb
}

func (UserDb) TableName() string {
	return "users"
}
func (UsersDb) TableName() string {
	return "users"
}
