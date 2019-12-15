# micro.demo

This example makes use of the micro.

## Usage

1、build \*.proto file
```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. rpc.proto
```
or
```
protoc --micro_out=. --go_out=. rpc.proto
```
2、Run the micro API with the rpc handler

```
micro api --handler=rpc
```
3、Run server

```
go run ingredient/main.go
```
4、Run api

```
go run rpc/mian.go
```

or usage
```
consul agent -dev -bind [local address]
micro --registry=consul api --handler=rpc
<project srv> --registry=consul
<project api> --registry=consul
```

Make a POST request to recipe/add which will call go.micro.api.recipe recipe.Add

```
curl -H 'Content-Type: application/json' -d '{"name": "dog", "ingredient":{"name":"dog"}}' "http://localhost:8080/recipe/add"
```

Make a POST request to /recipe/GetRecipeByName which will call go.micro.api.recipe recipe.GetRecipeByName

```
curl -H 'Content-Type: application/json' -d '{"name": "dog"}' "http://localhost:8080/recipe/GetRecipeByName"
```

## Note

Micro doesn't run very well when it starts under windows. It often fails to find a service. After finding a service, it is basically stable.