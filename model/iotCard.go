package model

import "github.com/jinzhu/gorm"

type IotCard struct {
	gorm.Model     //不要这个执行操作自动创建结构体字段数据库表，执行再打开这个就可以获取增删改的操作时间
	Id         int `json:"id"`
	//Name       string `json:"name"`    //收件人
	//Mobile     string `json:"mobile"`  //电话
	Address string `json:"address"` //地址
	Number  string `json:"number"`  //出售出的卡号
	Courier string `json:"courier"` //快递
	Weixin  string `json:"weixin"`  //微信
	Remarks string `json:"remarks"` //备注

}
