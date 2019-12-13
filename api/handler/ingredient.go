package handler

import (
	"context"
	"fmt"

	"github.com/nalaws/micro.demo/srv/ingredient/proto/ingredient"
)

type IngredientApi struct {
	Client ingredient.IngredientService
}

func (api *IngredientApi) Add(ctx context.Context, req *ingredient.NewIngredient, rsp *ingredient.Response) error {
	fmt.Println("@@@ Received IngredientApi.Add API request")

	// make the request
	response, err := api.Client.Add(ctx, req)
	if err != nil {
		return err
	}

	// set api response
	rsp.Code = response.Code
	rsp.Message = response.Message
	return nil
}
