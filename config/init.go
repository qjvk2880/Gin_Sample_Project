package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type configs struct {
	Service ServiceConfig `yaml:"service"`
	MySql   MySqlConfig   `yaml:"mysql"`
}

var DB *gorm.DB

var Configs configs

func Init(Config, ConfigPath *string) {
	var configPath string
	if Config == nil || *Config == "dev" {
		_, b, _, _ := runtime.Caller(0)
		BasePath := filepath.Dir(b)
		configPath = BasePath + "/file/configs.yaml"
	} else {
		configPath = *ConfigPath
	}
	load(configPath)
}

func load(ConfigsPath string) {
	yamlFile, err := os.ReadFile(ConfigsPath)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func DbUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Configs.MySql.User,
		Configs.MySql.Password,
		Configs.MySql.Host,
		Configs.MySql.Port,
		Configs.MySql.DBName,
	)
}
