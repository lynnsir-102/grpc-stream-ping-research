package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"tenyears102/grpc-stream-ping-research/grpc/pb"
)

const tag = "########## [research-server] ##########"

var port = flag.Int("port", 60020, "port number")

var kasp = keepalive.ServerParameters{
	/*
		如果一个 client 空闲超过设置时间, 发送一个 GOAWAY, 为了防止同一时间发送大量 GOAWAY, 会在设置时间间隔上下浮动 t*10%, 即 t+10%t 或者 t-10%t
		If a client is idle for set seconds, send a GOAWAY
		MaxConnectionIdle:     15 * time.Second,
	*/
	/*
		如果任意连接存活时间超过设置时间, 发送一个 GOAWAY
		If any connection is alive for more than set seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second,
	*/
	/*
		在强制关闭连接之间, 允许有设置的时间完成 pending 的 rpc 请求
		If set below 1s, a minimum value of 1s will be used instead.
		Allow set seconds for pending RPCs to complete before forcibly closing connections
		MaxConnectionAgeGrace: 5 * time.Second,
	*/

	/*
		如果一个 client 空闲超过设置时间, 则发送一个 ping 请求
		The current default value is 2 hours
		Ping the client if it is idle for set seconds to ensure the connection is still active
	*/
	Time: 10 * time.Second, // niffler server Time 10s
	/*
		如果 ping 请求在设置时间内未收到回复, 则认为该连接已断开
		The current default value is 20 seconds.
		Wait set second for the ping ack before assuming the connection is dead
	*/
	Timeout: 3 * time.Second,
}

var kaep = keepalive.EnforcementPolicy{
	/*
		如果客户端两次 ping 的间隔小于设置时间, 则关闭连接
		The current default value is 5 minutes
		If a client pings more than once every set seconds, terminate the connection
	*/
	MinTime: 10 * time.Second, // niffler server MinTime 10s
	/*
		如果没有 active 的 stream, 是否允许发送 ping
		Allow pings even when there are no active streams
	*/
	PermitWithoutStream: true,
}

func main() {
	flag.Parse()

	sigchan := make(chan os.Signal)
	eventchan := make(chan string)

	go func() {
		signal.Notify(sigchan, syscall.SIGUSR1)

		for sig := range sigchan {
			fmt.Printf("%s recv sig: %s\n\n", tag, sig.String())
			eventchan <- time.Now().String()
		}
	}()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("%s failed to listen: %v", tag, err)
	}

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	pb.RegisterServerSideDeliveryServer(s, &service{ch: eventchan})

	fmt.Printf("%s begin serve on pid %d\n\n", tag, os.Getpid())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("%s failed to serve: %v", tag, err)
	}
}

type service struct {
	ch <-chan string
}

func (s *service) Watch(req *pb.ServerSideDeliveryRequest, srv pb.ServerSideDelivery_WatchServer) error {
	for {
		select {
		case <-srv.Context().Done():
			fmt.Printf("%s client close\n\n", tag)
			return nil
		case key := <-s.ch:
			err := srv.Send(&pb.ServerSideDeliveryResponse{Key: key})
			if err == nil {
				fmt.Printf("%s server send data finish, data: %s\n\n", tag, key)
				continue
			}
			fmt.Printf("%s server send data fail, data: %s, err: %s\n\n", tag, key, err.Error())
		}
	}
}
