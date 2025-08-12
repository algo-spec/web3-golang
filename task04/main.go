package main

import (
	"log"
	"strconv"
	"task04/app/routes"
	"task04/domain/database"
	"task04/domain/models"
	"task04/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置
	if _, err := database.Connect(); err != nil {
		log.Printf("[ERROR] 连接数据库失败: %v", err)
		return
	}
	defer database.Close()

	log.Printf("[INFO] 连接数据库成功")

	//创建表
	models.AutoMigrate(database.GetDB())

	//启动 Gin
	r := gin.Default()

	//注册路由
	routes.RegisterRoutes(r)

	config.LoadConfig()
	r.Run(":" + strconv.Itoa(config.AppConfig.ServerPort))
}
