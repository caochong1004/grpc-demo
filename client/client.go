package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "github.com/longjoy/grpc-demo/proto"
	"log"
)

var port string

func init()  {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}


func main()  {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())

	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	//一元rpc
	_ = SayHello(client)

	//服务端流式rpc
	//m := new(pb.HelloRequest)
	//_ = SayList(client, m)

	//客户端流式rpc
	//m := new(pb.HelloRequest)
	//_ = SayRecord(client, m)

	//双向流式rpc
	//m := new(pb.HelloRequest)
	//_ = SayRoute(client, m)
}

//一元rpc
func SayHello(client pb.GreeterClient) error  {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name:"caochong"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

//服务端流式rpc
//func SayList(client pb.GreeterClient, r *pb.HelloRequest) error  {
//	stream, _ := client.SayList(context.Background(), r)
//	for  {
//		resp, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			return err
//		}
//		log.Printf("resp: %v", resp)
//	}
//	return nil
//}

//客户端流式rpc
//func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error  {
//	stream, _ := client.SayRecord(context.Background())
//	for n := 0; n <= 6; n++ {
//		_ = stream.Send(r)
//	}
//	resp, _ := stream.CloseAndRecv()
//	log.Printf("resp err: %v", resp)
//	return nil
//}

//双向流式rpc
//func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error  {
//	stream, _ := client.SayRoute(context.Background())
//	for n := 0; n <= 6; n++{
//		_ = stream.Send(r)
//		resp, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			return err
//		}
//		log.Printf("resp: %v", resp)
//
//	}
//	_ = stream.CloseSend()
//	return nil
//}
