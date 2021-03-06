package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"example.com/faster/faster/fasterpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Starting client...")
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("34.233.106.96:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %w", err)
		return
	}
	defer conn.Close()

	c := fasterpb.NewFasterServiceClient(conn)

	getNetworkSpeed(c)
}

func getNetworkSpeed(c fasterpb.FasterServiceClient) {
	fmt.Println("Infering network speed\n")

	amounts := [4]int{
		//maybe take into account bytes issue by gRPC when encoding messages
		//see notes.md
		int(math.Pow(2, 17)),
		int(math.Pow(2, 18)),
		int(math.Pow(2, 19)),
		int(math.Pow(2, 20)),
	}

	for _, amount := range amounts {
		fmt.Printf("Sending %dkb\n", amount/1000)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		sendPayload(ctx, c, make([]byte, amount))
		cancel()
	}

}

func sendPayload(ctx context.Context, c fasterpb.FasterServiceClient, payload []byte) (int, bool) {
	payloadWeight := len(payload)

	_, err := c.InferNetworkSpeed(ctx, &fasterpb.InferNetworkSpeedRequest{
		Payload: payload,
	})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Printf("Error: %v\n\n", statusErr.Message())
				return payloadWeight, false
			} else {
				fmt.Printf("unexpected error: %v", statusErr.Code())
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
		return payloadWeight, false
	}
	fmt.Printf("Network can handle %dkb per second of traffic\n", payloadWeight/1000)
	return payloadWeight, true
}
