package controllor

import (
	"net/http"

	"example.com/m/v2/model"
	"example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)

//获取拼多多弹出层的添加数据页面
func EjectPdd(c *gin.Context) {
	c.HTML(200, "ejectPdd.html", nil)
}

//查询拼多多所有客户信息数据
func PddQuery(c *gin.Context) {
	db := mysql.Dbinit()
	var pdd []model.PddData
	db.Find(&pdd)
	c.JSON(200, pdd)
}

//添加拼多多客户信息数据
func PddAddData(c *gin.Context) {
	db := mysql.Dbinit()
	//Name := c.PostForm("name")       //姓名
	//Mobile := c.PostForm("mobile")   //电话
	// Source
	// file, _ := c.FormFile("file")    //上传文件的
	Address := c.PostForm("address") //地址
	ComCode := c.PostForm("comcode") //商品编码
	Courier := c.PostForm("courier") //快递
	Price := c.PostForm("price")     //出售价
	Cost := c.PostForm("cost")       //成本
	Profit := c.PostForm("profit")   //利润
	Oder := c.PostForm("oder")       //订单编号
	Remarks := c.PostForm("remarks") //备注
	newClass := model.PddData{
		//Name:    Name,
		//Mobile:  Mobile,

		Address: Address,
		ComCode: ComCode,
		Courier: Courier,
		Price:   Price,
		Cost:    Cost,
		Profit:  Profit,
		// Profit:  Profit,
		Oder:    Oder,
		Remarks: Remarks,
	}

	db.Create(&newClass)
	c.JSON(200, "提交成功")
}

// 删除拼多多客户信息
func DeletePdd(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	var delete model.PddData
	db.Where(id).Delete(&delete) //调试前端时添加了id
	db.Delete(id)
	// c.Redirect(http.StatusMovedPermanently, "/queryadd") //重定向
	c.JSON(200, "删除成功")
}

// 修改拼多多数据
//修改数据  //实现了指定字段修改内容
func Updatapdd(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	//Name := c.PostForm("name")       //姓名
	//Mobile := c.PostForm("mobile")   //电话
	Address := c.PostForm("address") //地址
	ComCode := c.PostForm("comcode") //商品编码
	Courier := c.PostForm("courier") //快递
	Price := c.PostForm("price")     //出售价
	Cost := c.PostForm("cost")       //成本
	Profit := c.PostForm("profit")   //利润
	Oder := c.PostForm("oder")       //订单编号
	Remarks := c.PostForm("remarks") //备注

	//这些没问题的你可以去接口那个试一下
	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段 //更新有变化的值，没有变化的不更新
	db.Model(&model.PddData{}).Where(id).Update(&model.PddData{
		//Name:    Name,
		//Mobile:  Mobile,
		Address: Address,
		ComCode: ComCode,
		Courier: Courier,
		Price:   Price,
		Cost:    Cost,
		Profit:  Profit,
		Oder:    Oder,
		Remarks: Remarks,
	})
	// db.Model(&model.Article{}).Where(id).Update(&model.Article{Title: title, Content: content})
	// db.Model(&user).Where(id).Updates(User{Name: Name, Mobile: Mobile, Add: Add, Bz: Bz})
	c.JSON(200, "修改成功")
	c.Redirect(http.StatusMovedPermanently, "/getpageadd") //重定向

}

// 获取修改弹出层页面
func Updatepage(c *gin.Context) {
	c.HTML(200, "updatepdd.html", nil)
}

// 根据订单编号查询数据
// 通过备注信息进行模糊搜索
func OderQ(c *gin.Context) {
	db := mysql.Dbinit()
	var pdddata []model.PddData
	//id := c.Param("id")
	// Name := c.PostForm("name")       //姓名
	// Mobile := c.PostForm("mobile")   //电话
	//Address := c.Query("address") //地址
	//Number := c.Query("number") //出售的卡号信息
	//Courier := c.Query("courier") //快递
	// Weixin := c.PostForm("weixin")   //客户的微信
	Oder := c.Query("oder") //备注
	// 传一个参数完全匹配查询出来的
	//db.Where("number =  ?", Number).Debug().Find(&iot)
	// 实际模糊搜索功能
	db.Where("oder LIKE  ?", Oder+"%").Find(&pdddata)
	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)

	c.JSON(200, pdddata)

}

//通过地址搜索
func AddQ(c *gin.Context) {
	db := mysql.Dbinit()
	var pdddata []model.PddData
	//id := c.Param("id")
	// Name := c.PostForm("name")       //姓名
	// Mobile := c.PostForm("mobile")   //电话
	//Address := c.Query("address") //地址
	//Number := c.Query("number") //出售的卡号信息
	//Courier := c.Query("courier") //快递
	// Weixin := c.PostForm("weixin")   //客户的微信
	Address := c.Query("address") //备注
	// 传一个参数完全匹配查询出来的
	//db.Where("number =  ?", Number).Debug().Find(&iot)
	// 实际模糊搜索功能
	db.Where("address LIKE  ?", Address+"%").Find(&pdddata)
	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)

	c.JSON(200, pdddata)

}
