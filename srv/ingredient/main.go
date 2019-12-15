package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"

	"github.com/nalaws/micro.demo/srv/ingredient/handler"
	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.ingredient"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	ingredient.RegisterIngredientHandler(service.Server(), new(handler.Ingredient))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
