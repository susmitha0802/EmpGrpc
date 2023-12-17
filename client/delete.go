package main

import (
	"context"
	pb "emp_directory/proto"
	"log"
	"time"
)

func callSayDeleteEmp(client pb.EmpServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "Rithika"
	res, err := client.SayDeleteEmp(ctx, &pb.DeleteEmpRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not Delete: %v", err)
	}

	log.Println("Employee with name:", name, res.GetMessage())
}
