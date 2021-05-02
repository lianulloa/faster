package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"example.com/faster/faster/fasterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	fasterpb.UnimplementedFasterServiceServer
}

func (*server) InferNetworkSpeed(ctx context.Context, req *fasterpb.InferNetworkSpeedRequest) (*fasterpb.InferNetworkSpeedResponse, error) {
	fmt.Println("InferNetworkSpeed function was invoked")

	if ctx.Err() == context.Canceled {
		msg := fmt.Sprintf("Network cannot handle %dkb\n", len(req.Payload)/1000)
		fmt.Println("Canceled")
		fmt.Println(msg)
		return nil, status.Error(
			codes.Canceled,
			msg,
		)
	}
	res := &fasterpb.InferNetworkSpeedResponse{
		Speed: fmt.Sprint((req.Payload)),
	}
	return res, nil
}

func main() {
	fmt.Println("Starting faster...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	fasterpb.RegisterFasterServiceServer(s, &server{})

	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		} else {
			fmt.Println("Server started")

		}

	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program")
}
