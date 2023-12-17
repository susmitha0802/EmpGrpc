package main

import (
	"context"
	pb "emp_directory/proto"
	"log"
	"time"
)

func callSayUpdateEmp(client pb.EmpServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "Amrutha"
	res, err := client.SayUpdateEmp(ctx, &pb.UpdateEmpRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not Update: %v", err)
	}

	log.Println("Employee with name:", name, res.GetMessage())
}
