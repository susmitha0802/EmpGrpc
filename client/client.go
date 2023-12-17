package main

import (
	"log"

	pb "emp_directory/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect the server %v", err)
	}

	defer conn.Close()

	client := pb.NewEmpServiceClient(conn)
	callSayCreateEmp(client)
	callSayReadEmp(client)
	callSayUpdateEmp(client)
	callSayDeleteEmp(client)
}
