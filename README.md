# micro.demo

This example makes use of the micro.

## Usage

### 1、build \*.proto file
build all \*.proto file. file path: api\rpc\proto、srv\ingredient\proto\ingredient、srv\recipe\proto\recipe.

e.g:
```
protoc --micro_out=. --go_out=. rpc.proto
```
### 2、Run micro

```
micro api --handler=rpc --address=0.0.0.0:8080
```
```
micro api --handler=api --address=0.0.0.0:8081
```
```
micro web
```
### 3、Run server

```
go run srv/ingredient/main.go
```
```
go run srv/recipe/main.go
```
### 4、Run api

```
go run api/api/main.go
```
```
go run api/rpc/main.go
```
### 5、Run web

```
go run web/main.go
```
### 6、Test
```
curl -H 'Content-Type: application/json' -d '{"name": "dog"}' "http://localhost:8080/recipe/GetRecipeByName"
```
```
curl http://localhost:8081/recipe/GetRecipeByName?name=dog
```
Open with browser: `http://localhost:8082/recipe`

## Note

Micro doesn't run very well when it starts under windows. It often fails to find a service. After finding a service, it is basically stable.
