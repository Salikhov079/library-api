package handler

import pb "github.com/Salikhov079/library-api/genprotos"

type Handler struct {
	Author   pb.AuthorServiceClient
	Book     pb.BookServiceClient
	Genre    pb.GenreServiceClient
	Borrower pb.BorrowerServiceClient
}

func NewHandler(
	author pb.AuthorServiceClient, book pb.BookServiceClient, 
	genre pb.GenreServiceClient, borrower pb.BorrowerServiceClient,
) *Handler {
	return &Handler{author, book, genre, borrower}
}
