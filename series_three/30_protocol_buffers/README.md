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

