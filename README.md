# go-rpc
Demonstrate GoLang RPC
* Protocol buffers are ideal when you want the communication between microservices fast and efficient
* With protocol buffers you are using protobuf instead of JSON and this makes the transported payload much smaller and the Marshall/UnMarshalling is much faster
* With gRPC, the interaction between client and server is via a procedure rather than a JSON data as in typical REST based services

## Context
* The context that is passed by the client to the server is quite important
  * In  this case, we are just passing context.Background() which is just a process running in background
  * Important information such as timeouts and other signals can be shared betwwen client and server using this
  * for example, look at the following code
    * Here we are sending a cancellation signal from client to server and make the server cancel the connection when there is a value  in the channel
    * client and server communicate by means of a channel
    
```go
Client Side:
===========
// A client can create a conext in the following manner
ctx, cancel := context.WithCancel(context.Background())
// Here client's defer function would be invoked when the app is closed on the client side
// This passes on a message in the ctx channel
defer cancel()

Server Side:
===========
//Server can listen on the ctx channel and return error when a message is found on the channel

select {
case <-ctx.Done() :
  return nil, ctx.Err()
}
```

## ProtoBuf code generation
* The .proto file is created with the request and response type
* Also, the remote function that is invoked is mentioned as a service
* The following command generates the code for service(stub) and setters/getters for the request, response as in normal protocol buffer files

```go
protoc -I=. --go_out=plugins=grpc:. <name of the .proto file>

plugins=grpc is responsible for generating the stub code for service in the pb.go file
```
* NOTE: only protobuf 3 is allowed for gRPC
* NOTE: You have to export the variables on the ".proto" file for them to be available outside their package 

## How to use
* start the server by executing "go run rpc_server.go"
* This will start a server on port 12345 on localhost
* Invoke the client by executing go run rpc_client.go
  * You will see the message printed on stdout. This is after the procedure was executed on the server side
