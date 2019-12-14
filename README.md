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

Make a POST request to /example/call which will call go.micro.api.example Example.Call

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example/call"
```

Make a POST request to /example/foo/bar which will call go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
