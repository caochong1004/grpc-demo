package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/longjoy/grpc-demo/proto"
	"net"
)

var port string

func init()  {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct {}

//一元rpc
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error){
	return &pb.HelloReply{Message:"Hello World"}, nil
}

//服务端流式rpc
//func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer ) error  {
//	for n := 0; n <= 6; n++ {
//		_ = stream.Send(&pb.HelloReply{Message:"hello.list"})
//	}
//	return nil
//}

//客户端流式rpc
//func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error  {
//	for  {
//		resp, err := stream.Recv()
//		if err == io.EOF {
//			message := &pb.HelloReply{Message:"say.record"}
//			return stream.SendAndClose(message)
//		}
//		if err != nil {
//			return err
//		}
//		log.Printf("resp: %v", resp)
//	}
//	return nil
//}

//双向流式rpc
//func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error  {
//	n := 0
//	for  {
//		_ = stream.Send(&pb.HelloReply{Message:"say.route"})
//		resp, err := stream.Recv()
//		if err == io.EOF {
//			return nil
//		}
//		if err != nil {
//			return err
//		}
//		n++
//		log.Printf("resp: %v", resp)
//	}
//}

func main()  {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	fmt.Println("监听地址 %s\n",port)
	server.Serve(lis)
}
