package router

import (
	"example.com/m/v2/controllor"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	//router.LoadHTMLGlob("view/*")

	router.LoadHTMLGlob("view/*")
	router.Static("/layui", "./layui")

	router.GET("/", controllor.Index)
	router.GET("/ejectpdd", controllor.EjectPdd)        //获取拼多多弹出层添加的页面
	router.POST("/pddadddata", controllor.PddAddData)   //添加拼多多客户信息数据
	router.GET("/pddquery", controllor.PddQuery)        //查询拼多多数据
	router.GET("/deletepdd/:id", controllor.DeletePdd)  //执行删除拼多多客户信息
	router.POST("/updatepdd/:id", controllor.Updatapdd) //执行修改拼多多客户信息
	router.GET("/updatepage", controllor.Updatepage)    //修改拼多多修改信息页面

	router.GET("/iotcard", controllor.IotCard)       //获取物联卡用户信息页面
	router.GET("/iotquery", controllor.IotQuery)     //查询物联卡客户信息的数据
	router.POST("/upiot/:id", controllor.UpIot)      //修改物联卡客户信息数据
	router.POST("/addiot", controllor.AddIot)        //添加物联卡用户信息数据
	router.GET("/getaddpage", controllor.GetAddPage) //获取添加物联卡用户信息的弹出层页面
	router.GET("/deleiot/:id", controllor.DeleIot)   //删除物联卡数据
	router.GET("/upIiotpage", controllor.UpIotPage)  //获取修改弹出物联卡数据信息页面

	router.GET("/login", controllor.LoginPage)            //获取登录页面的路由
	router.GET("/register", controllor.Register)          //获取注册页面的路由
	router.POST("/registeruser", controllor.RegisterUser) //注册用户信息

	router.Run(":8087")

}
