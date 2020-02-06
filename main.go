package main

import (
	"context"
	"fmt"
	_ "gin_blog/docs"
	"gin_blog/pkg/setting"
	"gin_blog/routers"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Swagger Example API
// @version 1.6.5
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1


func main()  {
	r := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		ReadHeaderTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server ...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	    if err := s.Shutdown(ctx); err != nil {
		    log.Fatal("Server Shutdown:", err)
		}
	log.Println("Server exiting")
}
