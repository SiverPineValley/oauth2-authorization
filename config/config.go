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
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
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
