package handler

import (
	pb "github.com/Salikhov079/library-api/genprotos"

	"github.com/gin-gonic/gin"
)

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
		panic(err)
	}

	_, err = h.Book.CreateBook(ctx, arr)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	_, err = h.Book.UpdateBook(ctx, &arr)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll Book
// @Description 	GetAll page
// @Tags			Book
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.BookReq true "Query parameter"
// @Success 		200  {object}  pb.AllBooks     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAlld"
// @Router 			/book/getall [get]
func (h *Handler) GetAllBooks(ctx *gin.Context) {
	book := &pb.BookReq{}
	book.Title = ctx.Param("title")

	res, err := h.Book.GetAllBooks(ctx, book)
	if err != nil {
		panic(err)
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
// @Failure 		401  {string}  string     "Error while GetByIdd"
// @Router 			/book/get/{id} [get]
func (h *Handler) GetByIdBook(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Book.GetByIdBook(ctx, &id)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, res)
}
