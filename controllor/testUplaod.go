package controllor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"example.com/m/v2/model"
	"example.com/m/v2/mysql"
	"github.com/gin-gonic/gin"
)

func GetUptestHtml(c *gin.Context) {
	c.HTML(200, "GetUptestHtml.html", nil)
}

// func TestUploFile(c *gin.Context) {
// 	form, _ := c.MultipartForm()

// 	fs := form.File["fs"]

// 	for index, f := range fs {
// 		log.Println(f.Filename)
// 		dst := fmt.Sprintf("./static/imgTest/%s-%d", f.Filename, index)
// 		// 上传文件到指定的目录
// 		c.SaveUploadedFile(f, dst)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("%d uploaded!", len(fs)),
// 	})

// }

// func TestUploFile(c *gin.Context) {

// 	//获取前端传来的数据
// 	form, _ := c.MultipartForm()
// 	// 获取所有图片
// 	files := form.File["fs"]
// 	// 遍历所有图片
// 	for _, file := range files {
// 		//获取图像后缀
// 		y := path.Ext(file.Filename)
// 		//获取时间戳防止重复 !需要精准到纳秒，防止传输过快产生同名，然后出错
// 		a := time.Now().UnixNano()
// 		//获取一个1w以内的随机数
// 		b := rand.Intn(10000)
// 		//将时间辍的类型转换
// 		z := strconv.FormatInt(a, 10)
// 		//将随机数转换类型
// 		x := strconv.FormatInt(int64(b), 10)
// 		//写入保存位置与自定义名称，并且带上文件自带后缀名
// 		//dst := path.Join("img", z+x+y)
// 		dst := fmt.Sprintf("./static/imgTest/", z+x+y)
// 		// 存储文件
// 		_ = c.SaveUploadedFile(file, dst)
// 	}
// 	//返回数据
// 	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))

// }

//多张图片上传  //自己在本地实现了多图片上传功能
// func TestUploFile(c *gin.Context) {
// 	c.Request.ParseMultipartForm(32 << 20)
// 	//获取所有上传文件信息
// 	fhs := c.Request.MultipartForm.File["fs"]

// 	uid := c.Request.FormValue("user_id")
// 	uploadDir := "./static/imgTest/" + uid
// 	err := os.MkdirAll(uploadDir, 0777)
// 	if err != nil {
// 		return
// 	}
// 	var i = 0
// 	for _, fheader := range fhs {
// 		i++
// 		newFileName := strconv.Itoa(i) + ".jpg" //这是修改上传上来文件的后缀名
// 		if err := c.SaveUploadedFile(fheader, uploadDir+"/"+newFileName); err != nil {
// 			//自己完成信息提示
// 			c.String(400, "图片上传失败")
// 			return
// 		}
// 	}
// 	c.String(200, "图片上传成功")
// }

//多张图片上传  //自己在本地实现了多图片上传功能  //实现了上传到数据了
// func TestUploFile(c *gin.Context) {
// 	db := mysql.Dbinit()
// 	name := c.PostForm("name") //商品名称
// 	test := c.PostForm("test") //商品编码
// 	c.Request.ParseMultipartForm(32 << 20)
// 	//获取所有上传文件信息
// 	fhs := c.Request.MultipartForm.File["fs"]

// 	uid := c.Request.FormValue("user_id")
// 	uploadDir := "./static/imgTest/" + uid
// 	err := os.MkdirAll(uploadDir, 0777)
// 	if err != nil {
// 		return
// 	}
// 	var i = 0
// 	for _, fheader := range fhs {
// 		i++
// 		newFileName := strconv.Itoa(i) + ".jpg" //这是修改上传上来文件的后缀名
// 		if err := c.SaveUploadedFile(fheader, uploadDir+"/"+newFileName); err != nil {
// 			//自己完成信息提示
// 			c.String(400, "图片上传失败")
// 			return
// 		}
// 		// new := model.ImgTest{Filepath: newFileName, Name: name, Test: test}
// 		new := model.ImgTest{Filepath: newFileName}
// 		db.Create(&new)
// 	}
// 	new := model.ImgTest{Name: name, Test: test}
// 	db.Create(&new)
// 	c.String(200, "图片上传成功")
// }

func TestUploFile(c *gin.Context) {
	db := mysql.Dbinit()
	name := c.PostForm("name") //测试标题
	//test := c.PostForm("test")
	new := model.ImgTest{Name: name}
	db.Create(&new)
	c.Request.ParseMultipartForm(32 << 20)
	//获取所有上传文件信息
	fhs := c.Request.MultipartForm.File["fs"]
	uid := c.Request.FormValue("user_id")
	uploadDir := "./static/img" + uid
	err := os.MkdirAll(uploadDir, 0777)
	if err != nil {
		return
	}
	var i = 0
	for _, fheader := range fhs {
		i++
		//newFileName := strconv.Itoa(i) + ".jpg" //这是修改上传上来文件的后缀名
		newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + ".jpg" //生成时间日期命名上传的文件
		if err := c.SaveUploadedFile(fheader, uploadDir+"/"+newFileName); err != nil {
			//自己完成信息提示
			c.String(400, "图片上传失败")
			return
		}
		// new := model.ImgTest{Filepath: newFileName, Name: name, Test: test}
		new := model.ImgTest{Filepath: newFileName}
		db.Create(&new)
	}

	c.String(200, "图片上传成功")
}

// 查询商品所有数据
func ImgQuery(c *gin.Context) {
	db := mysql.Dbinit()

	var ImgTest []model.ImgTest
	//var TestTitle []model.Commtidy
	db.Find(&ImgTest)
	//db.Find(&TestTitle)
	c.HTML(200, "demo.html", nil)
	//c.JSON(200, TestTitle)
	// c.JSON(200, ImgTest)

	defer db.Close()

}

//测试查询文章系统做示例的
func Article(c *gin.Context) {
	db := mysql.Dbinit()

	var ImgTest []model.ImgTest
	db.Find(&ImgTest)
	c.HTML(200, "TestArticle.html", ImgTest)
	defer db.Close()
}

//抄别人的
const ueBasePath = `static/upload/`

func ShowDemo(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the demo.html template
		"demo.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "UEditor演示",
		},
	)
}

func Action(c *gin.Context) {
	action := c.Query("action")
	datePath := time.Now().Format("20060102") + `/`

	switch action {
	//自动读入配置文件，只要初始化UEditor即会发生
	case "config":
		jsonByte, _ := ioutil.ReadFile("static/ueditor/conf/config.json")
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
		jsonByte = re.ReplaceAll(jsonByte, []byte(""))
		c.Writer.Write(jsonByte)

	case "uploadimage":
		{
			//创建保存文件的目录，每天一个目录
			err := os.MkdirAll(ueBasePath+`images/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, header, err := c.Request.FormFile("upfile")
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			name := time.Now().Format("150405") + header.Filename
			out, err := os.Create(ueBasePath + `images/` + datePath + name)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			data, err := json.Marshal(map[string]string{
				"url":      fmt.Sprintf("%s", ueBasePath+`images/`+datePath+name), //保存后的文件路径
				"title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
				"original": header.Filename,                                       //原始文件名
				"state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			if err != nil {
				panic(err)
			}
			c.Writer.Write(data)
		}

	case "uploadvideo":
		{
			err := os.MkdirAll(ueBasePath+`video/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的视频
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, err := c.FormFile("upfile")
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			name := time.Now().Format("150405") + file.Filename
			path := ueBasePath + `video/` + datePath
			if err := c.SaveUploadedFile(file, path+name); err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"url":      fmt.Sprintf("%s", ueBasePath+`images/`+datePath+name), //保存后的文件路径
				"title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
				"original": file.Filename,                                         //原始文件名
				"state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			c.Writer.Write(data)
		}

	case "uploadfile":
		{
			//创建保存文件的目录，每天一个目录
			err := os.MkdirAll(ueBasePath+`files/`+datePath, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//保存上传的文件
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			file, header, err := c.Request.FormFile("upfile")
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}

			name := time.Now().Format("150405") + header.Filename
			out, err := os.Create(ueBasePath + `files/` + datePath + name)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("上传文件错误: %s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"url":      fmt.Sprintf("%s", ueBasePath+`images/`+datePath+name), //保存后的文件路径
				"title":    "",                                                    //文件描述，对图片来说在前端会添加到title属性上
				"original": header.Filename,                                       //原始文件名
				"state":    "SUCCESS",                                             //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
			})
			c.Writer.Write(data)
		}
	case "uploadscrawl":
		{
			path := ueBasePath + `scrawl/` + datePath
			name := time.Now().Format("150405.999999") + `.jpg`
			err := os.MkdirAll(path, 0777)
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("创建目录错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			//upfile为base64格式文件，转成图片保存
			upfile, _ := c.GetPostForm("upfile")
			upBytes, err := base64.StdEncoding.DecodeString(upfile) // + "_" + filename
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("获取图片错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			err = ioutil.WriteFile(path+name, upBytes, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			if err != nil {
				data, _ := json.Marshal(map[string]string{
					"state": fmt.Sprintf("保存涂鸦文件错误:%s", err.Error()),
				})
				c.Writer.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]string{
				"state":    "SUCCESS",
				"url":      `/` + path + name,
				"title":    `涂鸦`,
				"original": `涂鸦不见了`,
			})
			c.Writer.Write(data)
			return
		}
	}
}
