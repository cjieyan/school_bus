package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-admin/database"
	orm "go-admin/global"
	mycasbin "go-admin/pkg/casbin"
	"go-admin/pkg/logger"
	//"go-admin/models"
	//"go-admin/models/gorm"
	"go-admin/router"
	"go-admin/tools/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title go-admin API
// @version 1.0.1
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 74520518 进入技术交流群 请备注，谢谢！
// @license.name MIT
// @license.url https://github.com/wenjianzhang/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	configName := "config/settings.yml"


	config.Setup(configName)
	logger.Setup()
	//3. 初始化数据库链接
	database.Setup(config.DatabaseConfig.Driver)
	//4. 接口访问控制加载
	mycasbin.Setup()

	gin.SetMode(gin.DebugMode)
	//
	//err := gorm.AutoMigrate(orm.Eloquent)
	//if err != nil {
	//	log.Fatalln("数据库初始化失败 err: %v", err)
	//}
	//
	//if config.ApplicationConfig.IsInit {
	//	if err := models.InitDb(); err != nil {
	//		log.Fatal("数据库基础数据初始化失败！")
	//	} else {
	//		config.SetApplicationIsInit()
	//	}
	//}

	r := router.InitRouter()

	defer orm.Eloquent.Close()

	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server Run ", config.ApplicationConfig.Host+":"+config.ApplicationConfig.Port)
	log.Println("Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

//func main() {
//	cmd.Execute()
//}
