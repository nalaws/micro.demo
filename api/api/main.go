package main

import (
	"log"

	"github.com/nalaws/micro.demo/api/api/handler"

	ingredientSrv "github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
	recipeSrv "github.com/nalaws/micro.demo/srv/recipe/proto/recipe"

	"github.com/micro/go-micro"
)

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.recipe"), // go.micro.api.*: * is impletement server name
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Recipe{
				Recipe:     recipeSrv.NewRecipeService("go.micro.srv.recipe", service.Client()),
				Ingredient: ingredientSrv.NewIngredientService("go.micro.srv.ingredient", service.Client()),
			},
		),
	)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
