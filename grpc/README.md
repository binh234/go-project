# Go-gRPC

This is an experiment project with gRPC using Go.

gRPC is a high-performance open-source Remote Procedure Call (RPC) framework originally developed by Google. It allows developers to build distributed systems and microservices by enabling client and server applications to communicate transparently using a simple interface definition language and a set of tools that generate efficient client and server stubs. It uses Protocol Buffers as the default serialization mechanism, which is a binary format that provides a compact and efficient way to encode structured data.

One of the main benefits of gRPC is its speed and efficiency, as it uses HTTP/2 as the underlying transport protocol and supports multiplexing, flow control, and other optimizations to reduce latency and bandwidth usage. Additionally, gRPC provides features such as bi-directional streaming, flow control, and deadline-based request cancellation.

gRPC is commonly used in developing microservices and blockchain applications. In microservices architecture, gRPC's high performance and low latency make it an excellent choice for inter-service communication, allowing developers to build scalable and efficient microservices-based systems. Additionally, gRPC's support for multiple programming languages and platforms makes it an ideal choice for building polyglot microservices systems.

In blockchain applications, gRPC can be used to implement the communication layer between different nodes in a blockchain network. It can also be used to connect blockchain nodes to other services and applications, such as data storage and analytics tools. gRPC's speed and efficiency make it a good choice for blockchain applications that require fast and reliable communication between different components.

### Preresquires
- [**Protocol buffer**](https://developers.google.com/protocol-buffers) compiler, `protoc`, [version 3](https://protobuf.dev/programming-guides/proto3)
  For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
- [**Go plugins**](https://grpc.io/docs/languages/go/quickstart/) for the protocol compiler:
  For installation instructions, see [Quickstart with gRPC in Go](https://grpc.io/docs/languages/go/quickstart/)

### Usage

The `client` and `server` folder implements diffrent types of communication between client and server (unary/client streaming/server streaming/bidirectional). Change the function call in `client/main.go` file to change the communication type.

You may need to recompile the updated `.proto` file to regenerate gRPC code for Go, run this command:

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

Open 2 separate terminals, one for server and another for client. Then run:

```bash
# Inside server folder
go run .
```

```bash
# Inside client folder
go run .
```