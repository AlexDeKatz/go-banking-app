package main

import (
	"github.com/AlexDeKatz/banking/app"
	"github.com/AlexDeKatz/banking/logging"
)

func main() {

	logging.Info("Starting the banking application...")
	app.Start()
}
