package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Log        Log        `json:"log"`
	HttpListen HttpListen `json:"http_listen"`
	DBConfig   DBConfig   `json:"db_config"`
	GrpcServer GrpcServer `json:"grpc_server"`
	Ampq       Ampq       `json:"ampq"`
	Scheduler  Scheduler  `json:"scheduler"`
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

type GrpcServer struct {
	Host string `json:"ip"`
	Port string `json:"port"`
}

type Ampq struct {
	Host     string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Scheduler struct {
	Period     string `json:"period"`
	BeforeTime string `json:"before_time"`
	EventTime  string `json:"event_time"`
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
		GrpcServer: GrpcServer{Host: viper.GetString("grpc.host"), Port: viper.GetString("grpc.port")},
		Ampq: Ampq{Host: viper.GetString("ampq.host"), Port: viper.GetString("ampq.port"),
			User: viper.GetString("ampq.user"), Password: viper.GetString("ampq.password")},
		Scheduler: Scheduler{Period: viper.GetString("scheduler.period"), BeforeTime: viper.GetString("scheduler.before_time"), EventTime: viper.GetString("scheduler.event_time")},
	}
}
