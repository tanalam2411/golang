
> 1. Inside your base go_grpc directory create directory go/src

```bash
go_grpc$ mkdir -p go/src
go_grpc$ protoc pb/messages.proto --go_out=plugins=grpc:go/src
go_grpc$ tree
├── go
│   └── src
│       └── github.com
│           └── ps
│               └── hellogrpc
│                   └── messages
│                       └── messages.pb.go
├── pb
│   └── messages.proto
└── README.md
```
