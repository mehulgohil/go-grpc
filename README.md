# grpc

## Introduction

### What is gRPC
gRPC is a robust open-source RPC (Remote Procedure Call) framework used to build scalable and fast APIs.
It allows the client and server applications to communicate transparently and develop connected systems

In simple terms, gRPC is a way for computers to talk to each other efficiently and reliably,
even if they speak different languages or run on different platforms.
It acts as a translator, making sure that messages are understood correctly and delivered safely.

Many leading tech firms have adopted gRPC, such as Google, Netflix, Square, IBM, Cisco, & Dropbox.
This framework relies on HTTP/2, protocol buffers,
and other modern technology stacks to ensure maximum API security, performance, and scalability.

In 2015,
Google developed gRPC as an extension of the RPC framework
to link many microservices created with different technologies.

## Evolution of RPC to gRPC
RPC stands for Remote Procedure Call.
As the name suggests, the idea is that we can invoke a function/method on a remote server.
RPC protocol allows one to get the result for a problem in the same format regardless of where it is executed.
It can be local or in a remote server using better resources.

The idea is the same.
An API is built by defining public methods.
Then the methods are called with arguments.
RPC is just a bunch of functions,
but in the context of an HTTP API,
it entails putting the method in the URL and the arguments in the query string or body.

gRPC is an adaptation of traditional RPC frameworks.
The most important difference is
that gRPC uses protocol buffers as the interface definition language for serialization and communication instead of JSON/XML.
Protocol buffers can describe the structure of data, 
and the code can be generated from that description for generating or parsing a stream of bytes
that represents the structured data.
This is the reason gRPC is preferred for the web applications that are polyglot
(implemented with different technologies).

```gRPC = RPC + Protobuff```

### RPC vs gRPC
* RPC is a generic protocol for remote procedure calls, while gRPC is a specific implementation of RPC that uses the HTTP/2 protocol for communication
* Also, RPC uses a binary encoding format to transmit data, while gRPC supports several serialization formats, including Protocol Buffers, JSON, and XML.
* gRPC is generally faster and more efficient than RPC, due to its use of HTTP/2 and the binary serialization format.
* gRPC has some additional features that are not available in traditional RPC, such as bidirectional streaming and server-side streaming.

## Architecture

![gRPCArch.png](.attachments%2FgRPCArch.png)

We have the gRPC client and server sides.
In gRPC, every client service includes a stub (auto-generated files),
similar to an interface containing the current remote procedures.
The gRPC client makes the local procedure call to the stub with parameters to be sent to the server.
The client stub then serializes the parameters with the marshaling process using Protobuf
and forwards the request to the local client-time library in the local machine.

The OS makes a call to the remote server machine via HTTP/2 protocol.
The server’s OS receives the packets and calls the server stub procedure,
which decodes the received parameters and executes the respective procedure invocation using Protobuf.
The server stub then sends back the encoded response to the client transport layer.
The client stub gets back the result message and unpacks the returned parameters,
and the execution returns to the caller.