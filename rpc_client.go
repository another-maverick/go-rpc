package main

import (
	"fmt"
	pb "github.com/another-maverick/go-rpc/rpcProtocolBuffer"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"os"
)

func main() {
	addr :=  "localhost:12345"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("cannot connect to the server on port 12345 ", err)
		os.Exit(1)
	}
	defer conn.Close()
	rpcClient := pb.NewTestRpcClient(conn)
	name := "Steph Curry"

	req := &pb.RpcRequest{Name: name}

	resp, err := rpcClient.TestRpc(context.Background(), req)
	if err != nil {
		fmt.Println("Cannot invoke the RPC method ", err)
		os.Exit(1)
	}
	fmt.Println(resp.Response)

}