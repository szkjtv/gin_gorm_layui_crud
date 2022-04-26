package model

import "github.com/jinzhu/gorm"

type PddData struct {
	gorm.Model     //不要这个执行操作自动创建结构体字段数据库表，执行再打开这个就可以获取增删改的操作时间
	Id         int `json:"id,omitempty"`
	//Name       string `json:"name"`    //收件人
	//Mobile     string `json:"mobile"`  //电话
	Address string `json:"address"` //地址
	ComCode string `json:"comcode"` //商品编码
	Courier string `json:"courier"` //快递
	Price   string `json:"price"`   //价格
	Cost    string `json:"cost"`    //成本
	Profit  string `json:"profit"`  //利润
	Oder    string `json:"oder"`    //订单编号
	Remarks string `json:"remarks"` //备注
	// CreateAt time.Time
	// UpdateAt time.Time
}
