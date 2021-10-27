package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"tenyears102/grpc-stream-ping-research/grpc/pb"
)

const tag = "########## [research-client] ##########"

var addr = flag.String("addr", "localhost:60020", "the address to connect to")

var kacp = keepalive.ClientParameters{
	/*
		如果没有 activity, 则每隔设置时间发送一个ping包
		If set below 10s, a minimum value of 10s will be used instead
		send pings every set seconds if there is no activity
	*/
	Time: 10 * time.Second, // rainbow client Time 10s
	/*
		如果 ping ack 设置时间之内未返回则认为连接已断开
		wait set second for ping ack before considering the connection dead
	*/
	Timeout: 3 * time.Second,
	/*
		如果没有 active 的 stream, 是否允许发送 ping
		send pings even without active streams
	*/
	PermitWithoutStream: true,
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("%s did not connect: %v", tag, err)
	}
	defer conn.Close()

	req := pb.ServerSideDeliveryRequest{Key: "hello world"}

CHECK:
	for {
		state := conn.GetState()
		if state == connectivity.Ready {
			break
		}
		fmt.Printf("%s client conn state change, state: %s\n\n", tag, state.String())
		conn.WaitForStateChange(context.TODO(), state)
	}

	client := pb.NewServerSideDeliveryClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Watch(ctx, &req)
	if err != nil {
		cancel()
		if stream != nil {
			_ = stream.CloseSend()
		}
		goto CHECK
	}

	fmt.Printf("%s client begin recv\n\n", tag)

	for {
		res, err := stream.Recv()
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				log.Printf("%s get stream recv failed, state: %v, err: %s\n\n", tag, s.Code(), err.Error())
				cancel()
				_ = stream.CloseSend()
				goto CHECK
			}
			log.Printf("%s get stream recv failed, unexpected err: %s\n\n", tag, err.Error())
			continue
		}
		log.Printf("%s get server response: %s\n\n", tag, res.Key)
	}
}
