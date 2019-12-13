package handler

import (
	"context"
	"fmt"

	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
)

type Ingredient struct{}

func (ing *Ingredient) Add(ctx context.Context, in *ingredient.NewIngredient, out *ingredient.Response) error {
	fmt.Println("@@@ Add come in.")
	out.Code = 200
	out.Message = "success"
	return nil
}
