package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"spm-serv/controller"
	"spm-serv/core"
	"spm-serv/dao"
	"syscall"
	"time"
)

func beforeStartup() *core.Config{
	err := core.Global.Load()
	if err!=nil {
		core.Log.Panicln(err)
	}
	core.InitLog(&core.Global)
	dao.InitDao(&core.Global)
	return &core.Global
}

func startup(config *core.Config){
	controller.SetMode(config)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(controller.Logger())
	controller.Validator()
	controller.Router(engine)
	serv := &http.Server{Addr:":" + config.Server.Port, Handler: engine}

	go func(){
		if err := serv.ListenAndServe(); err!=nil && err!=http.ErrServerClosed {
			core.Log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	core.Log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func(){
		cancel()
		dao.CloseDb()
	}()
	if err := serv.Shutdown(ctx); err != nil {
		core.Log.Fatal("Server Shutdown: ", err)
	}
	core.Log.Println("Server exiting")
}