package main

import (
	"context"
	pb "emp_directory/proto"
	"log"
	"time"
)

func callSayReadEmp(client pb.EmpServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayReadEmp(ctx, &pb.ReadEmpRequest{})
	if err != nil {
		log.Fatalf("Could not fetch emp list: %v", err)
	}
	log.Println("Name\tDept_Name\tManager_Name")

	for _, v := range res.GetEl() {
		log.Printf("%s\t%s\t\t%s\n", v.Name, v.DeptName, v.ManagerName)
	}
}
