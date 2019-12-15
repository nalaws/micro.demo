package handler

import (
	"context"
	"fmt"

	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
)

type Ingredient struct{}

func (ing *Ingredient) Add(ctx context.Context, in *ingredient.NewIngredient, out *ingredient.Response) error {
	fmt.Println("@@@ ingredient Add come in.")
	out.Code = 200
	out.Message = "success"
	return nil
}

func (ing *Ingredient) GetIngredientByName(ctx context.Context, in *ingredient.IngredientName, out *ingredient.IngredientInfo) error {
	fmt.Println("@@@ ingredient GetIngredientByName come in.")
	out.Name = "success"
	return nil
}
