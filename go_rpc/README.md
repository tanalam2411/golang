

> 1 . Install gRPC tooling
```sh
go_rpc$ go get -u google.golang.org/grpc
```

> 2 . Install protocol buffer tooling
```sh
go_rpc$ go get -u github.com/golang/protobuf/proto
go_rpc$ go get -u github.com/golang/protobuf/protoc-gen-go

#Single line go get
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

>3 . protoc
```sh
protoc -I ./pb/messages.proto --go_out=plugins=grpc:./src
```