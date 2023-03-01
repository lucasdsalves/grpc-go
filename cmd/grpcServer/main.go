package main

import (
	"database/sql"
	"grpc-go/internal/database"
	"grpc-go/internal/pb"
	service "grpc-go/internal/services"
	"net"

	// _ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./db.sqlite")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}
}
