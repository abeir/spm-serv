package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	vaildate "spm-serv/controller/validate"
	"spm-serv/core"
	"time"
)

func Router(engine *gin.Engine){

	engine.GET("/ping", Ping)
	engine.POST("/publish", Publish)
	engine.GET("/search", Search)
	engine.GET("/info", Info)
	engine.GET("/upgrade", Upgrade)
	engine.GET("/download", Download)
}

func Validator(){
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("checkVersion", vaildate.CheckVersion)
		if err!=nil {
			panic(err)
		}
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