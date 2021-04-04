package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var configs *Config

type server struct {
	Process string `json:"process"`
	Port    string `json:"port"`
	Host    string `json:"host"`
}

type database struct {
	Engine   string `json:"engine"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SslMode  bool   `json:"sslmode"`
}

type Config struct {
	Server   server   `json:"server"`
	Database database `json:"database"`
}

func InitConfig() (err error) {
	file, err := os.Open("./config/config.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&configs); err != nil {
		log.Fatal(err)
	}

	return
}

func GetProcessName() string {
	return configs.Server.Process
}

func GetPort() string {
	return configs.Server.Port
}

func GetHost() string {
	return fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)
}

func GetDBEngine() string {
	return configs.Database.Engine
}

func GetDBInfo() string {
	sslmode := "disable"
	if configs.Database.SslMode {
		sslmode = "require"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", configs.Database.Host, configs.Database.User, configs.Database.Password, configs.Database.Dbname, configs.Database.Port, sslmode)
}
