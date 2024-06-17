package handler

import (
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-api/genprotos"

	"github.com/gin-gonic/gin"
)

type FilterBookF struct {
	AuthorName string
	Title      string
}

// @Summary 		Create Book
// @Description 	Create page
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.BookReq    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/book/create [post]
func (h *Handler) CreateBook(ctx *gin.Context) {
	arr := &pb.BookReq{}
	err := ctx.BindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.Book.CreateBook(ctx, arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Update Book
// @Description 	Update page
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Param     		id 		path    string   true    "Book ID"
// @Param   		Update  body    pb.BookRes  true    "Update"
// @Success 		200   {string}  string   "Update Successful"
// @Failure 		401   {string}  string   "Error while updated"
// @Router			 /book/update/{id} [put]
func (h *Handler) UpdateBook(ctx *gin.Context) {
	arr := pb.BookRes{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.Book.UpdateBook(ctx, &arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Delete Book
// @Description 	Delete page
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true    "Book ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/book/delete/{id} [delete]
func (h *Handler) DeleteBook(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Book.DeleteBook(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll Book
// @Description 	GetAll page
// @Tags			Book
// @Accept  		json
// @Produce  		json
// @Param 			query  query   FilterBookF true "Query parameter"
// @Success 		200  {object}  pb.AllBooks     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAll"
// @Router 			/book/getall [get]
func (h *Handler) GetAllBooks(ctx *gin.Context) {
	book := &pb.FilterBook{}
	ctx.ShouldBindQuery(&book)

	res, err := h.Book.GetAllBooks(ctx, book)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		GetById Book
// @Description 	GetById page
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string      true   "Book ID"
// @Success 		200  {object}  pb.BookRes "GetById Successful"
// @Failure 		401  {string}  string     "Error while GetById"
// @Router 			/book/get/{id} [get]
func (h *Handler) GetByIdBook(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Book.GetByIdBook(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		Overdue Book
// @Description 	Overdue page
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Success 		200  {object}  string     "Overdue Successful"
// @Failure 		401  {string}  string     "Error while Overdue"
// @Router 			/book/overdue [get]
func (h *Handler) OverdueBooks(ctx *gin.Context) {
	var count int
	books, err := h.Borrower.GetAllIdBorrowers(ctx, &pb.Void{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	for i := 0; i < len(books.Borrowers); i++ {
		t, err := time.Parse("2006-01-02T15:04:05Z07:00", books.Borrowers[i].ReturnDate)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		if t.Before(time.Now()) {
			count++
			h.Borrower.DeleteBorrower(ctx, &pb.ById{Id: books.Borrowers[i].Id})
		}
	}

	ctx.JSON(200, fmt.Sprintf("Success! %d borrower return", count))
}
