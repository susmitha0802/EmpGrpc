package main

import (
	"context"
	"log"

	pb "emp_directory/proto"
)

func (s *empServer) SayDeleteEmp(ctx context.Context, req *pb.DeleteEmpRequest) (*pb.DeleteEmpResponse, error) {

	log.Println("Delete employee request received")

	name := req.GetName()

	res := s.db.Where("name = ?", name).Delete(&Employee{})

	if res.RowsAffected == 0 {
		if res.Error == nil {
			return &pb.DeleteEmpResponse{
				Message: "Already deleted / No record found to delete",
			}, nil
		}

		log.Println("Error", res.Error)
		return &pb.DeleteEmpResponse{}, res.Error
	}

	return &pb.DeleteEmpResponse{
		Message: "Deleted successfully",
	}, nil
}
