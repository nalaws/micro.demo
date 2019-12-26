package main

import (
	"net/http"

	"github.com/nalaws/micro.demo/web/handler"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.recipe"),
		web.Version("latest"),
	)

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html"))) // http://localhost:8082/recipe

	// register call handler
	service.HandleFunc("/download", handler.Export) // http://localhost:8082/recipe/export

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
