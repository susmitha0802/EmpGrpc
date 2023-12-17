package main

import (
	"context"
	"log"

	pb "emp_directory/proto"
)

func (s *empServer) SayUpdateEmp(ctx context.Context, req *pb.UpdateEmpRequest) (*pb.UpdateEmpResponse, error) {

	log.Println("Update employee request received")

	name := req.GetName()

	res := s.db.Model(&Employee{}).Where("name = ?", name).Update("name", name+" Varshini")


	if res.RowsAffected == 0 {
		if res.Error == nil {
			return &pb.UpdateEmpResponse{
				Message: "Already updated / No record found to update",
			}, nil
		}
		log.Println("Error", res.Error)
		return &pb.UpdateEmpResponse{}, res.Error
	}

	return &pb.UpdateEmpResponse{
		Message: "Updated successfully",
	}, nil
}
