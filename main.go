package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"spm-serv/controller"
	"spm-serv/core"
	"syscall"
	"time"
)

func beforeStartup(){
	err := core.Global.Load()
	if err!=nil {
		panic(err)
	}
}

func startup(){
	engine := gin.Default()
	controller.Validator()
	controller.Router(engine)
	serv := &http.Server{Addr:":8000", Handler: engine}

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
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		core.Log.Fatal("Server Shutdown: ", err)
	}
	core.Log.Println("Server exiting")
}

func main() {
	beforeStartup()
	startup()
}