package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hossainemruz/hello-grpc/calculator"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("failed to connect with the server")
	}
	defer conn.Close()

	// create calculator client
	client := calculator.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Add two number
	fmt.Println("========== Adding 50 and 38 ===============")
	resp, err := client.Add(ctx, &calculator.Operands{A: 50, B: 38})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result: ", resp.Sum)

	// sum a list of numbers
	fmt.Println("=========== Summing: 1,2,3,4,5,6,7,8,9 =============")
	resp, err = client.Sum(ctx, &calculator.Numbers{Num: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result: ", resp.Sum)

	fmt.Println("Closing client......")
}
