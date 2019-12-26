package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	recipeSrv "github.com/nalaws/micro.demo/srv/recipe/proto/recipe"
)

func Export(w http.ResponseWriter, r *http.Request) {
	fmt.Println("@@@ Received Download web request")
	// call the backend service
	cli := recipeSrv.NewRecipeService("go.micro.srv.recipe", client.DefaultClient)
	rsp, err := cli.GetRecipeByName(context.Background(), &recipeSrv.RecipeName{Name: "dog"})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// return case 1 or case 2
	// 1: download file
	/*	filePath := "xxxx"
		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "文件名.扩展名"))
		w.Header().Add("Content-Type", "application/octet-stream")
		//http.ServeFile(w, r, filePath)
		fileData, err := ioutil.ReadFile(des)
		w.Write(fileData)
		// delete temp file
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("failed to remove temp: %v\n", err)
		}
		return
	*/

	//2: we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Name,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
