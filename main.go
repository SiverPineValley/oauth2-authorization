package main

import (
	"log"
	"oauth2-authorization/config"
	"oauth2-authorization/middlewares"
	"oauth2-authorization/router"
	"oauth2-authorization/utility"
	"runtime"

	"github.com/urfave/negroni"
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

	// 3. Init Router
	router := router.InitRouter()

	// 4. Request Logger Middlewares
	n := negroni.Classic()
	n.UseHandler(router)
	n.UseHandler(middlewares.RequestLogger())

	n.Run(":" + config.GetPort())
}
