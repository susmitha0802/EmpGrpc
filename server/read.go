package main

import (
	"context"
	pb "emp_directory/proto"
	"log"
)

func (s *empServer) SayReadEmp(ctx context.Context, req *pb.ReadEmpRequest) (*pb.ReadEmpResponse, error) {

	log.Println("Fetch employee list received")

	e := []*pb.EmpList{}
	s.db.Debug().Table("employees").
		Select("employees.id, employees.name, departments.name as dept_name, managers.name as manager_name").
		Joins("join departments on employees.dept_id = departments.id").
		Joins("join employees as managers on employees.manager_id = managers.id").
		Where("employees.deleted_at IS NULL").
		Find(&e)

	return &pb.ReadEmpResponse{
		El: e,
	}, nil
}
