package handler

import (
	"context"
	"fmt"

	api "github.com/nalaws/micro.demo/api/rpc/proto/recipe"
	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
	"github.com/nalaws/micro.demo/srv/recipe/proto/recipe"
)

type RecipeApi struct {
	Recipe     recipe.RecipeService
	Ingredient ingredient.IngredientService
}

func (r *RecipeApi) Add(ctx context.Context, in *api.NewRecipe, out *api.Response) error {
	fmt.Println("@@@ Received RecipeApi.Add RPC request")

	// make the request
	_, err := r.Recipe.Add(ctx, &recipe.NewRecipe{
		Name: in.Name,
	})
	ingredientRsp, err := r.Ingredient.Add(ctx, &ingredient.NewIngredient{
		Name:    in.Name,
		Species: ingredient.Species_Animal,
	})
	if err != nil {
		return err
	}

	// set api response
	out.Code = ingredientRsp.Code
	out.Message = ingredientRsp.Message
	return nil
}

func (r *RecipeApi) GetRecipeByName(ctx context.Context, in *api.RecipeName, out *api.Response) error {
	fmt.Println("@@@ Received RecipeApi.GetRecipeByName RPC request")

	// make the request
	recipeRsp, err := r.Recipe.GetRecipeByName(ctx, &recipe.RecipeName{
		Name: in.Name,
	})
	ingredientRsp, err := r.Ingredient.GetIngredientByName(ctx, &ingredient.IngredientName{
		Name: in.Name,
	})
	if err != nil {
		return err
	}

	// set api response
	out.Code = 200
	out.Message = "{" + recipeRsp.GetName() + "," + ingredientRsp.GetName() + "}"
	return nil
}
