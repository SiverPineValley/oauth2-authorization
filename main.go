package main

import (
	"log"
	"oauth2-authorization/config"
	"oauth2-authorization/middlewares"
	"oauth2-authorization/router"
	"oauth2-authorization/utility"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 1. Init Config
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Init Logger
	utility.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Init Database
	err = middlewares.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Init Router
	e := router.InitRouter()

	// 4. Request Logger Middlewares
	e.Use(middlewares.RequestLogger)

	e.Start(":" + config.GetPort())
}
