package main

import (
	"log"



	"github.com/lollmark/calculator_go/internal/application"

)

func main() {
	app := application.New()
	log.Fatal(app.RunServer())
}
