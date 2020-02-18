# hello-grpc
A sample gRPC server and client

## Resources

- [Protocol Buffer](https://developers.google.com/protocol-buffers/docs/proto3)
- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [gRPC Concepts](https://www.grpc.io/docs/guides/)
- [gRPC Basics - Go](https://www.grpc.io/docs/tutorials/basic/go/)
- [gRPC Go Examples](https://github.com/grpc/grpc-go/tree/master/examples)
- [Mocking Service for gRPC](https://github.com/grpc/grpc-go/blob/master/Documentation/gomock-example.md)

## Pre-requisite

- Install dependencies:
    ```console
    sudo apt-get update
    sudo apt-get install curl unzip git build-essential automake libtool
    ```
- Install protocol buffer v3 compiler:
    ```console
    git clone git@github.com:protocolbuffers/protobuf.git
    cd protobuf/
    git checkout v3.11.4
    ./autogen.sh
    ./configure
    make
    sudo make install
    sudo ldconfig
    ```
 - Verify protoc verion:
  ```console
   protoc --version
   libprotoc 3.11.4
  ```
- Install protoc Go Plugin:
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Generate Code

Run the commands in project root directory

```console
protoc --go_out=plugins=grpc:. calculator/*.proto
```

## Implement Server

## Write Client


## Test gRPC server

- Run server:
```console
$ go run ./server/main.go
Adding.... A: 50 B: 38
Summing a list of numbers.......
```

- Run Client:
```console
$ go run ./client/main.go
========== Adding 50 and 38 ===============
Result:  88
=========== Summing: 1,2,3,4,5,6,7,8,9 =============
Result:  45
Closing client......
```