# RPC API

This example makes use of the micro API with the RPC handler.

The api in this mode lets you serve http while routing to standard go-micro services via RPC.

The micro api with rpc handler only supports POST method and expects content-type of application/json or application/protobuf.

## Usage

编译proto文件
```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. rpc.proto
```
或者
```
protoc --micro_out=. --go_out=. rpc.proto
```

Run the micro API with the rpc handler

```
micro api --handler=rpc
```

Run this example

```
go run rpc.go
```

Make a POST request to /example/call which will call go.micro.api.example Example.Call

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example/call"
```

Make a POST request to /example/foo/bar which will call go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
