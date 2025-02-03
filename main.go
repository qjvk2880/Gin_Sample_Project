package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/model"
	httpEngine "example/web-service-gin/router/http"
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	config.Init(configFlag, prodConfigPath)

	config.DB, err = gorm.Open(mysql.Open(config.DbUrl()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return
	}

	fmt.Println("0000")
	err := config.DB.AutoMigrate(&model.Album{}, &model.Review{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("0000")

	httpEngine.Run(config.Configs.Service.HttpPort)
}
