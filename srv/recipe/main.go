package main

import (
	"log"

	"github.com/micro/go-micro"

	"github.com/nalaws/micro.demo/srv/recipe/handler"
	"github.com/nalaws/micro.demo/srv/recipe/proto/recipe"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.recipe"),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	recipe.RegisterRecipeHandler(service.Server(), new(handler.Recipe))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
