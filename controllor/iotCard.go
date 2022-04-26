package controllor

import (
	"net/http"

	"example.com/m/v2/model"
	"example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)

//获取页面
func IotCard(c *gin.Context) {
	c.HTML(200, "iotpage.html", nil)
}

// post方法添加数据
func AddIot(c *gin.Context) {
	db := mysql.Dbinit()
	//Name := c.PostForm("name")       //姓名
	//Mobile := c.PostForm("mobile")   //电话
	Address := c.PostForm("address") //地址
	Number := c.PostForm("number")   //出售的卡号信息
	Courier := c.PostForm("courier") //快递
	Weixin := c.PostForm("weixin")   //客户的微信
	Remarks := c.PostForm("remarks") //备注
	newClass := model.IotCard{
		//Name:    Name,
		//Mobile:  Mobile,
		Address: Address,
		Number:  Number,
		Courier: Courier,
		Weixin:  Weixin,
		Remarks: Remarks,
	}

	db.Create(&newClass)
	c.JSON(200, "提交成功")
}

// 修改物联卡客户信息
func UpIot(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	//Name := c.PostForm("name")       //姓名
	//Mobile := c.PostForm("mobile")   //电话
	Address := c.PostForm("address") //地址
	Number := c.PostForm("number")   //出售的卡号信息
	Courier := c.PostForm("courier") //快递
	Weixin := c.PostForm("weixin")   //客户的微信
	Remarks := c.PostForm("remarks") //备注
	db.Model(&model.IotCard{}).Where(id).Update(&model.IotCard{
		//Name:    Name,
		//Mobile:  Mobile,
		Address: Address,
		Number:  Number,
		Courier: Courier,
		Weixin:  Weixin,
		Remarks: Remarks,
	})
	// db.Model(&model.Article{}).Where(id).Update(&model.Article{Title: title, Content: content})
	// db.Model(&user).Where(id).Updates(User{Name: Name, Mobile: Mobile, Add: Add, Bz: Bz})
	c.JSON(200, "修改物联卡信息成功")
	c.Redirect(http.StatusMovedPermanently, "/iotcard") //重定向

	// c.JSON(200, "提交成功")
}

// 查询物联卡客户信息的数据
func IotQuery(c *gin.Context) {
	db := mysql.Dbinit()
	var iot []model.IotCard
	db.Find(&iot)
	//c.HTML(200, "pddquery.html", pdd)
	c.JSON(200, iot)
	// defer db.Close()
}

// 获取添加数据的弹出层页面
func GetAddPage(c *gin.Context) {
	c.HTML(200, "iotAdd.html", nil)
}

// 删除数据
func DeleIot(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	var delete model.IotCard
	db.Where(id).Delete(&delete) //调试前端时添加了id
	db.Delete(id)
	// c.Redirect(http.StatusMovedPermanently, "/queryadd") //重定向
	c.JSON(200, "删除成功")
}

// 获取弹出层的修改页面 物联卡客户信息
func UpIotPage(c *gin.Context) {
	c.HTML(200, "UpIotPage.html", nil)
}
