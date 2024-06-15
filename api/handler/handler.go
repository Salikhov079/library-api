package handler

import pb "github.com/Salikhov079/library-api/genprotos"

type Handler struct {
	Author   pb.AuthorServiceClient
	User     pb.UserServiceClient
	Book     pb.BookServiceClient
	Genre    pb.GenreServiceClient
	Borrower pb.BorrowerServiceClient
}

func NewHandler(
	author pb.AuthorServiceClient, user pb.UserServiceClient, book pb.BookServiceClient,
	genre pb.GenreServiceClient, borrower pb.BorrowerServiceClient,
) *Handler {
	return &Handler{author, user,  book, genre, borrower}
}
