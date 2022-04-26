package router

import (
	"example.com/m/v2/controllor"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	//router.LoadHTMLGlob("view/*")

	router.LoadHTMLGlob("view/*")
	//router.Static("/imgagess", "./imgagess")
	router.Static("/static", "./static")

	// 拼多多下单用户地址信息数据增删改查
	router.GET("/", controllor.Index)
	router.GET("/ejectpdd", controllor.EjectPdd)        //获取拼多多弹出层添加的页面
	router.POST("/pddadddata", controllor.PddAddData)   //添加拼多多客户信息数据
	router.GET("/pddquery", controllor.PddQuery)        //查询拼多多数据
	router.GET("/deletepdd/:id", controllor.DeletePdd)  //执行删除拼多多客户信息
	router.POST("/updatepdd/:id", controllor.Updatapdd) //执行修改拼多多客户信息
	router.GET("/updatepage", controllor.Updatepage)    //修改拼多多修改信息页面

	//物联卡用户信息的增删改查
	router.GET("/iotcard", controllor.IotCard)       //获取物联卡用户信息页面
	router.GET("/iotquery", controllor.IotQuery)     //查询物联卡客户信息的数据
	router.POST("/upiot/:id", controllor.UpIot)      //修改物联卡客户信息数据
	router.POST("/addiot", controllor.AddIot)        //添加物联卡用户信息数据
	router.GET("/getaddpage", controllor.GetAddPage) //获取添加物联卡用户信息的弹出层页面
	router.GET("/deleiot/:id", controllor.DeleIot)   //删除物联卡数据
	router.GET("/upIiotpage", controllor.UpIotPage)  //获取修改弹出物联卡数据信息页面
	router.GET("/oderq", controllor.OderQ)           //根据订单编号查询数据
	router.GET("/addq", controllor.AddQ)             //根据地址来查询

	// 登录和注册的路由
	router.GET("/login", controllor.LoginPage)            //获取登录页面的路由
	router.GET("/register", controllor.Register)          //获取注册页面的路由
	router.POST("/registeruser", controllor.RegisterUser) //注册用户信息
	router.GET("/queryadd", controllor.GetAddree)         //根据号码字段查询数据
	router.GET("/addreeq", controllor.Addree)             //通过地址模糊查询数据
	router.GET("/remarksq", controllor.Remarks)           //通过备注信息进行模糊搜索

	// 炊大皇商品数据增删改查
	router.POST("/commodity", controllor.Commodity)        //上传单文件图片拉接口
	router.POST("/commpost", controllor.CommPost)          //上传商品信息表单数据和图片文件分开一个接口写
	router.GET("/ommindex", controllor.CommIndex)          //获取商品编码页面
	router.GET("/commadd", controllor.CommAdd)             //获取添加商品编码弹出层
	router.GET("/quercomm", controllor.QuerComm)           //查询商品所有数据
	router.GET("/spcode", controllor.Spcode)               //通过商品编码搜索查询数据
	router.GET("/delec/:id", controllor.DeleC)             //通过id来删除商品数据信息
	router.POST("/upcomm/:id", controllor.UpComm)          //通过id来修改内容
	router.GET("/getPagespname", controllor.GetPageSpname) //获取炊大皇的商品编辑页面

	router.Run(":8087")

}
