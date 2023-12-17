package main

import (
	"fmt"
	"log"
	"net"

	pb "emp_directory/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	port = ":8080"
)

type empServer struct {
	pb.EmpServiceServer
	db *gorm.DB
}

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=susmitha password=gorm dbname=emp_directory",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	db.Migrator().DropTable(&Department{}, &Employee{})
	db.AutoMigrate(&Department{}, &Employee{})

	depts := []Department{
		{ID: 1, Name: "Air"},
		{ID: 2, Name: "Pro"},
		{ID: 3, Name: "Pact"},
	}
	db.Model(&Department{}).Create(&depts)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEmpServiceServer(grpcServer, &empServer{
		db: db,
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the grpc server %v", err)
	}
}

type Department struct {
	ID   int
	Name string
}

type Employee struct {
	gorm.Model
	Name      string `gorm:"unique"`
	ManagerID int
	DeptID    int
	Dept      Department `gorm:"foreignKey:DeptID"`
}
