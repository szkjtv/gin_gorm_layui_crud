package controllor

import (
	"net/http"

	"example.com/m/v2/model"
	"example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)

// 获取登录页面
func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 获取注册页面
func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

//注册账号
func RegisterUser(c *gin.Context) {
	db := mysql.Dbinit()
	acount := c.PostForm("account") //地址
	pas := c.PostForm("password")   //出售的卡号信息
	city := c.PostForm("city")      //快递
	sex := c.PostForm("sex")        //客户的微信

	newClass := model.RgisterUser{
		//Name:    Name,
		//Mobile:  Mobile,
		Acount: acount,
		Pas:    pas,
		City:   city,
		Sex:    sex,
	}
	db.Create(&newClass)
	// c.JSON(200, "注册成功")
	c.Redirect(http.StatusMovedPermanently, "/login") //重定向
}

func LoginCorim(c *gin.Context) {

}
