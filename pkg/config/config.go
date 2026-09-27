package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Upload   UploadConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port string
	Mode string
	Host string
}

type DatabaseConfig struct {
	Driver          string
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	Charset         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type JWTConfig struct {
	Secret string
	Expire string
	Issuer string
}

type UploadConfig struct {
	Path       string
	MaxSize    int64
	AllowTypes []string
}

type LogConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
}

var AppConfig *Config

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败：%w", err)
	}

	viper.AutomaticEnv()

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败：%w", err)
	}

	AppConfig = config
	return config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.max_idle_conns", "10")
	viper.SetDefault("database.max_open_conns", "100")
}
