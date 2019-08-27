package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"spm-serv/controller/validation"
	"spm-serv/core"
	"time"
)

func Router(engine *gin.Engine){
	spmController := SpmController{}
	engine.GET("/ping", spmController.Ping)
	engine.POST("/publish", spmController.Publish)
	engine.GET("/search", spmController.Search)
	engine.GET("/info", spmController.Info)
	engine.GET("/upgrade", spmController.Upgrade)
	engine.GET("/download", spmController.Download)

	console := engine.Group("/console")
	{
		consoleController := ConsoleController{}
		console.POST("/login", consoleController.Login)
		console.POST("/logout", consoleController.Logout)
	}
	upgrade := engine.Group("/upgrade")
	{
		upgradeController := UpgradeVersionController{}
		upgrade.GET("/list", upgradeController.List)
		upgrade.GET("/info", upgradeController.Info)
		upgrade.PUT("/upload", upgradeController.Upload)
		upgrade.PUT("/release", upgradeController.Release)
		upgrade.PUT("/detain", upgradeController.Detain)
	}
	pkg := engine.Group("/package")
	{
		packageController := PackageController{}
		pkg.GET("/list", packageController.List)
		pkg.GET("/info", packageController.Info)
		pkg.PUT("/enable", packageController.Enable)
		pkg.PUT("/disable", packageController.Disable)
	}
}

func Validator(){
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.Register(v, "checkVersion", validation.CheckVersion)
		validation.Register(v, "checkValues", validation.CheckValues)
	}
}

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		core.Log.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func SetMode(config *core.Config){
	if config.IsDev() {
		gin.SetMode(gin.DebugMode)
	}else if config.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}else if config.IsTest() {
		gin.SetMode(gin.TestMode)
	}
}