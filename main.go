package main

import (
	"stephanusnugraha/udemy-rest-api-golang/app"
	"stephanusnugraha/udemy-rest-api-golang/logger"
)

func main() {

	//log.Println("starting our application...")
	logger.Info("Starting the application...")
	app.Start()
}
