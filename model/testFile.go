package model

import "github.com/jinzhu/gorm"

type ImgTest struct {
	gorm.Model //导出这个获取时间增删改字段
	// Id         int
	Name string
	//Test     string
	Filepath string `json:"filepath"` //上传文件后保存的路径
	//Filename string `json:"filename"` //上传文件后重新生成的名称
	//Createtime int64  `json:"createtime"`
	//CommtidyId int
	// Commtidy   Commtidy `json:"ac" gorm:"Commtidy"`
	//	CommtidyId int
	//	Commtidy `gorm:"foreignkey:Commtidy"`
	//Commtidy   []Commtidy

	// CommtidyID int //`json:"Commtidy_Id"`
	// Commtidy   Commtidy `gorm:"foreignkey:CommtidyId"` //指定关联外键
	//Commtidy Commtidy `gorm:"foreignkey:Commtidy:Spname"` //指定关联外键
	// CommtidyID int
	Commtidy Commtidy `gorm:"foreignKey:Spname"` //使用Commtidy表里的Spname作为外键
	// 使用 UserName 作为外键
}
