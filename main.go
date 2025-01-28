package main

import (
	"example/web-service-gin/config"
	httpEngine "example/web-service-gin/router/http"
	"flag"
)

func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	config.Init(configFlag, prodConfigPath)

	httpEngine.Run(config.Configs.Service.HttpPort)
}
