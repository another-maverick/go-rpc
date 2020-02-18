package main

import (
	"fmt"
	pb "github.com/another-maverick/go-rpc/rpcProtocolBuffer"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"net"
)

type server struct {}

func (s *server) TestRpc(ctx context.Context, input *pb.RpcRequest) (*pb.RpcResponse, error){
	msg := "Hello " + input.Name
	return &pb.RpcResponse{Response: msg}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("cannot make the sever listen on 12345 ", err)
	}
	serv := grpc.NewServer()
	pb.RegisterTestRpcServer(serv, &server{})
	serv.Serve(ln)
}
