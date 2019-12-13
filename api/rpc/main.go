package main

import (
	"log"

	"github.com/nalaws/micro.demo/api/handler"
	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"

	"github.com/micro/go-micro"
)

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.ingredient"), // ingredient 能进， ingredientapi 进不来
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	ingredient.RegisterIngredientHandler(service.Server(), &handler.IngredientApi{
		// Create Service Client
		Client: ingredient.NewIngredientService("go.micro.srv.ingredient", service.Client()),
	})

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
