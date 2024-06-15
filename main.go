package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/library-api/api"
	"github.com/Salikhov079/library-api/api/handler"
	"github.com/Salikhov079/library-api/config"
	pb "github.com/Salikhov079/library-api/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
	}

	connU, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
	}
	
	cnf := config.Load()
	book := pb.NewBookServiceClient(conn)
	user := pb.NewUserServiceClient(connU)
	author := pb.NewAuthorServiceClient(conn)
	genre := pb.NewGenreServiceClient(conn)
	borrower := pb.NewBorrowerServiceClient(conn)

	handler := handler.NewHandler(author, user, book, genre, borrower)
	r := api.NewGin(handler)

	fmt.Println("Server listening on port:8081")
	err = r.Run(cnf.HTTPPort)
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
