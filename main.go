package main

import (
	"context"
	"fmt"
	"time"

	pb "test_go/price"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051" // Address of your Rust gRPC server
	queryPeriod = time.Second
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to dial server: %v", err)
		return
	}
	defer conn.Close()

	// Create a client instance using the connection.
	client := pb.NewPriceServiceClient(conn)

	// Loop to query the server every second.
	for {
		queryPriceData(client)
		time.Sleep(queryPeriod)
	}
}

// Function to query price data from the server.
func queryPriceData(client pb.PriceServiceClient) {
	// Context for the gRPC request.
	ctx := context.Background()

	// Making a gRPC request to get price data.
	response, err := client.GetPriceData(ctx, &pb.PriceDataRequest{Id: "BTC"})
	if err != nil {
		fmt.Printf("Error querying price data: %v\n", err)
		return
	}

	// Print the received price data.
	for _, priceData := range response.PriceDataList {
		fmt.Printf("Received price data: %v\n", priceData)
	}
}
