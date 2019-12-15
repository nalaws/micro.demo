package handler

import (
	"context"
	"fmt"

	"github.com/nalaws/micro.demo/srv/recipe/proto/recipe"
)

type Recipe struct{}

func (r *Recipe) Add(ctx context.Context, in *recipe.NewRecipe, out *recipe.Response) error {
	fmt.Println("@@@ recipe Add come in.")
	out.Code = 200
	out.Message = "success"
	return nil
}

func (r *Recipe) GetRecipeByName(ctx context.Context, in *recipe.RecipeName, out *recipe.RecipeInfo) error {
	fmt.Println("@@@ recipe GetRecipeByName come in.")
	out.Name = "success"
	return nil
}
