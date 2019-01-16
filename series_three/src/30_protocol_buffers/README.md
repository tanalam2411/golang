### Protocol Buffers -  are a language-neutral, platform-neutral extensible mechanism for serializing structured data. https://developers.google.com/protocol-buffers/

Installing Protocol Buffer(protoc) - Ubuntu 16.04

```sh
$ sudo apt-get install autoconf automake libtool curl make g++ unzip
```

Download from here - https://github.com/protocolbuffers/protobuf/releases

```sh
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```
or

```sh
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz
$ tar -xzvf protobuf-all-3.6.1.tar.gz
$ ./configure 
$ make
$ make check
$ sudo make install
$ sudo ldconfig
$ protoc --version
```

Generate todo.pb.go
```
series_three/src/30_protocol_buffers/todo$ protoc --go_out=. todo.proto
```

Run
```sh
series_three/src/30_protocol_buffers$ go run cmd/todo/main.go add hello there

series_three/src/30_protocol_buffers$ ls -1
cmd
mydb.pb
README.md
todo
```

Check file type, hexdump and pb raw data from mydb.pb
```sh
series_three/src/30_protocol_buffers$ file mydb.pb
mydb.pd: ASCII text

series_three/src/30_protocol_buffers$ hexdump -c mydb.pb
0000000  \n  \v   h   e   l   l   o       t   h   e   r   e
000000d

series_three/src/30_protocol_buffers$ cat mydb.pb | protoc --decode_raw
1: "hello there"
```


```
export GOPATH=$GOPATH:/home/tan/tanveer/golang/series_three
```