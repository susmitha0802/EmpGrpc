package main

import (
	"context"
	pb "emp_directory/proto"
	"log"
	"time"
)

func callSayCreateEmp(client pb.EmpServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	emp := pb.EmpDetails{Name: "Rithika", DeptId: 1, ManagerId: 1}

	res, err := client.SayCreateEmp(ctx, &pb.CreateEmpRequest{E: &emp})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	if res.GetE().Id == 0 {
		log.Fatalf("Not created successfully: %v", err)
	}

	log.Printf("Employee with id %v, name %v created successfully", res.GetE().Id, res.GetE().Name)
}
