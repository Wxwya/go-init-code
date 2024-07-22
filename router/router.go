package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"xwya/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://2d41-117-147-77-4.ngrok-free.app", "http://127.0.0.1:8000"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	initStaticRouter(r)
	router := r.Group("/api")
	initUserRouter(router)
	initDictionariesRouter(router)
	initProjectRouter(router)
	initFormRouter(router)
	server := &http.Server{
		Addr:    utils.Prot,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err, "启动失败")
			return
		}
	}()
	<-ctx.Done()
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err, "关闭失败")
		return
	}

}
