package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	vaildate "spm-serv/controller/validate"
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