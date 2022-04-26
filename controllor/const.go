package controllor

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"example.com/m/v2/model"
	"example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)

//定义一个创建文件目录的方法
// func Mkdir(basePath string) string {
// 	//	1.获取当前时间,并且格式化时间
// 	folderName := time.Now().Format("2006/01/02")
// 	folderPath := filepath.Join(basePath, folderName)
// 	//使用mkdirall会创建多层级目录
// 	os.MkdirAll(folderPath, os.ModePerm)
// 	return folderPath
// }

// 提交商品表单数据
func CommPost(c *gin.Context) {
	db := mysql.Dbinit()
	spname := c.PostForm("sname")    //商品名称
	spcode := c.PostForm("spcode")   //商品编码
	spconst := c.PostForm("spconst") //商品成本
	new := model.Commtidy{Spname: spname, Spcode: spcode, Spconst: spconst}
	db.Create(&new)
	c.JSON(200, "添加商品数据成功")
	defer db.Close()
}

// 上传商品编码，商品名称，图片，成本价格
// layui 前端上传文件只能使用一个接口，文件和表单数据要分开写
func Commodity(c *gin.Context) {

	db := mysql.Dbinit()
	spname := c.PostForm("spname")   //商品名称
	spcode := c.PostForm("spcode")   //商品编码
	spconst := c.PostForm("spconst") //商品成本
	fm, err := c.FormFile("file")    //上传图片文件的

	if err != nil {
		// responseErr(c, err)
		fmt.Print("错误信息请检查")
		return
	}

	if err != nil {
		//responseErr(c, err)
		return
	}

	//自动生成年月日路径
	now := time.Now()
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fm.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" || fileExt == ".mp4" {
		fileType = "img"
	}

	//文件存放路径
	fileDir := fmt.Sprintf("static/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		//responseErr(c, err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fm.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	c.SaveUploadedFile(fm, filePathStr)

	if fileType == "img" {
		new := model.Commtidy{Spname: spname, Spcode: spcode, Spconst: spconst, Filepath: filePathStr, Filename: fileName} //Filepath: filePathStr, Filename: fileName, Createtime: timeStamp
		// new := model.Commtidy{}
		//new := model.Album{0, filePathStr, fileName, 0, timeStamp}

		db.Create(&new)
		defer db.Close()
	}
	c.JSON(200, gin.H{
		"data": filePathStr,
	})
	//c.JSON(200, "添加商品数据成功")

}

// 获取商品编码页面
func CommIndex(c *gin.Context) {
	c.HTML(200, "commindex.html", nil)
}

// 获取商品编码的弹出层添加数据
func CommAdd(c *gin.Context) {
	c.HTML(200, "commadd.html", nil)
}

// 查询商品所有数据
func QuerComm(c *gin.Context) {
	db := mysql.Dbinit()
	var comm []model.Commtidy
	db.Find(&comm)
	//c.HTML(200, "pddquery.html", pdd)
	c.JSON(200, comm)
	// defer db.Close()

}

// 通过条件搜索查询
func Spcode(c *gin.Context) {
	db := mysql.Dbinit()
	var spcode []model.Commtidy
	Spcode := c.Query("spcode") //地址
	//Spname := c.Query("spname")                          //地址
	db.Where("spcode LIKE  ?", Spcode+"%").Find(&spcode) //原来的
	//db.Where("spcode LIKE  ?", Spcode+"%", "Spname LIKE ?", Spname+"%").Find(&spcode)
	// db.Where("spcode LIKE  ?", Spcode+"%").Order("Spname LIKE ?", Spname+"%").Find(&spcode)
	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)

	c.JSON(200, spcode)

}

// 通过id删除删除数据
func DeleC(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	var delete model.Commtidy
	db.Where(id).Delete(&delete) //调试前端时添加了id
	db.Delete(id)
	// c.Redirect(http.StatusMovedPermanently, "/queryadd") //重定向
	c.JSON(200, "删除成功")
}

// 更新商品信息
func UpComm(c *gin.Context) {
	db := mysql.Dbinit()
	id := c.Param("id")
	spname := c.PostForm("spname")   //商品名称
	spcode := c.PostForm("spcode")   //商品编码
	spconst := c.PostForm("spconst") //商品成本
	fm, err := c.FormFile("file")    //上传图片文件的
	if err != nil {
		// responseErr(c, err)
		fmt.Print("错误信息请检查")
		return
	}

	if err != nil {
		//responseErr(c, err)
		return
	}

	//自动生成年月日路径
	now := time.Now()
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fm.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" || fileExt == ".mp4" {
		fileType = "img"
	}

	//文件存放路径
	fileDir := fmt.Sprintf("static/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		//responseErr(c, err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fm.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	c.SaveUploadedFile(fm, filePathStr)

	db.Model(&model.Commtidy{}).Where(id).Update(&model.Commtidy{
		//Name:    Name,
		//Mobile:  Mobile,
		Spname:   spname,
		Spcode:   spcode,
		Spconst:  spconst,
		Filepath: filePathStr,
		Filename: fileName,
	})
	// db.Model(&model.Article{}).Where(id).Update(&model.Article{Title: title, Content: content})
	// db.Model(&user).Where(id).Updates(User{Name: Name, Mobile: Mobile, Add: Add, Bz: Bz})
	// c.JSON(200, "修改商品数据信息成功")
	//layUi前端提示图片请求接口异常返回一个这样的json格式的数据就可以了
	c.JSON(200, gin.H{
		"data": filePathStr,
	})
	//c.Redirect(http.StatusMovedPermanently, "/iotcard") //重定向

	// c.JSON(200, "提交成功")
}

// 获取炊大皇的商品编辑页面
func GetPageSpname(c *gin.Context) {
	c.HTML(200, "commup.html", nil)
}
