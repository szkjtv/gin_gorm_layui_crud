package model

import "github.com/jinzhu/gorm"

type Commtidy struct {
	gorm.Model //导出这个获取时间增删改字段
	Id         int
	//Params     string                `json:"params"`
	//File       *multipart.FileHeader `json:"file"`
	Spname   string //商品名称
	Spcode   string //商品编码
	Spconst  string //商品成本
	Filepath string `json:"filepath"` //上传文件后保存的路径
	Filename string `json:"filename"` //上传文件后重新生成的名称
	//Createtime int64  `json:"createtime"`
}
