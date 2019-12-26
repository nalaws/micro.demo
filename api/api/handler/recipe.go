package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	api "github.com/micro/go-micro/api/proto"
	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
	"github.com/nalaws/micro.demo/srv/recipe/proto/recipe"
)

// struct name is impletement grpc server name
type Recipe struct {
	Recipe     recipe.RecipeService
	Ingredient ingredient.IngredientService
}

type IngredientReq struct {
	Name string
}

type RecipeReq struct {
	Name       string
	Ingredient IngredientReq
}

// curl -H 'Content-Type: application/json' -d '{"name": "dog","ingredient": {"name": "dog"}}' http://localhost:8080/recipe/add
func (r *Recipe) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {
	fmt.Println("@@@ Received RecipeApi.Add API request")

	recipeReq, ingredientReq, err := r.validateAddParams(req)
	if err != nil {
		return nil
	}

	recipeRsp, err := r.Recipe.Add(ctx, recipeReq)
	if err != nil {
		return nil
	}
	ingredientRsp, err := r.Ingredient.Add(ctx, ingredientReq)
	if err != nil {
		return nil
	}

	// set api response
	rsp.Body = recipeRsp.Message + ingredientRsp.Message

	return nil
}

func (r *Recipe) validateAddParams(req *api.Request) (*recipe.NewRecipe, *ingredient.NewIngredient, error) {
	// check method
	if req.Method != "POST" {
		return nil, nil, errors.New("Method error.")
	}

	var recipeReq RecipeReq
	if err := json.Unmarshal([]byte(req.Body), &recipeReq); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal validateAddParams: %v", err)
	}

	return &recipe.NewRecipe{Name: recipeReq.Name},
		&ingredient.NewIngredient{Name: recipeReq.Ingredient.Name, Species: ingredient.Species_Animal},
		nil
}

// curl http://localhost:8080/recipe/GetRecipeByName?name=dog
func (r *Recipe) GetRecipeByName(ctx context.Context, req *api.Request, rsp *api.Response) error {
	fmt.Println("@@@ Received RecipeApi.GetRecipeByName API request")

	recipeReq, ingredientReq, err := r.validateGetRecipeByNameParams(req)
	if err != nil {
		return nil
	}

	recipeRsp, err := r.Recipe.GetRecipeByName(ctx, recipeReq)
	if err != nil {
		return nil
	}
	ingredientRsp, err := r.Ingredient.GetIngredientByName(ctx, ingredientReq)
	if err != nil {
		return nil
	}

	// set api response
	rsp.Body = recipeRsp.Name + ingredientRsp.Name

	return nil
}

func (r *Recipe) validateGetRecipeByNameParams(req *api.Request) (*recipe.RecipeName, *ingredient.IngredientName, error) {
	// check method
	if req.Method != "GET" {
		return nil, nil, errors.New("Method error.")
	}

	if req.Get["name"] == nil || req.Get["name"].Values[0] == "" {
		return nil, nil, errors.New("Params error.")
	}

	return &recipe.RecipeName{Name: req.Get["name"].Values[0]},
		&ingredient.IngredientName{Name: req.Get["name"].Values[0]},
		nil
}
