package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Log        Log        `json:"log"`
	HttpListen HttpListen `json:"http_listen"`
	DBConfig   DBConfig   `json:"db_config"`
}

type Log struct {
	LogFile  string `json:"log_file"`
	LogLevel string `json:"log_level"`
}

type HttpListen struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func GetConfigFromFile(filePath string) *Config {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't read configuration file: %s", err.Error())
	}
	return &Config{Log: Log{LogFile: viper.GetString("log.log_file"), LogLevel: viper.GetString("log.log_level")},
		HttpListen: HttpListen{Ip: viper.GetString("http_listen.ip"), Port: viper.GetString("http_listen.port")},
		DBConfig: DBConfig{User: viper.GetString("db.user"), Password: viper.GetString("db.password"),
			Host: viper.GetString("db.host"), Port: viper.GetString("db.port"), Database: viper.GetString("db.database")},
	}
}
