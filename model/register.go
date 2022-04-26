package model

import "github.com/jinzhu/gorm"

type RgisterUser struct {
	gorm.Model //自动获取增删改查的日期
	Id         int
	Acount     string //账号
	Pas        string //密码
	City       string //城市
	Sex        string //性别
}
