package main

import (
	"context"
	"log"

	pb "emp_directory/proto"
)

func (s *empServer) SayCreateEmp(ctx context.Context, req *pb.CreateEmpRequest) (*pb.CreateEmpResponse, error) {
	emp := Employee{
		Name:      req.E.GetName(),
		DeptID:    int(req.E.GetDeptId()),
		ManagerID: int(req.E.GetManagerId()),
	}

	log.Println("Create employee request received")

	res := s.db.Model(&Employee{}).Create(&emp)

	if res.RowsAffected == 0 {
		log.Println("Error", res.Error)
		return &pb.CreateEmpResponse{}, res.Error
	}

	response := pb.EmpDetails{Id: int32(emp.ID), Name: emp.Name, DeptId: int32(emp.DeptID), ManagerId: int32(emp.ManagerID)}
	return &pb.CreateEmpResponse{
		E: &response,
	}, nil
}
