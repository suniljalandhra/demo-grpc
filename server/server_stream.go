package main

import (
	"log"
	"time"

	pb "github.com/suniljalandhra/grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with name: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}
