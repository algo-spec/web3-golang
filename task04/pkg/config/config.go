package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBPath     string
	DBType     string
	ServerPort int
	LogLevel   string
	JwtSecret  string
}

var AppConfig *Config

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("base")
	v.SetConfigType("yaml")
	v.AddConfigPath("./pkg/config")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取 base.yaml 配置文件失败: %w", err))
	}

	var cfg Config
	cfg.DBPath = v.GetString("database.url")
	cfg.DBType = v.GetString("database.driver")
	cfg.ServerPort = v.GetInt("server.port")
	cfg.LogLevel = v.GetString("logging.level")
	cfg.JwtSecret = v.GetString("jwt.secret")

	AppConfig = &cfg
}
