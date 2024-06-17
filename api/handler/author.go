package handler

import (
	pb "github.com/Salikhov079/library-api/genprotos"

	"github.com/gin-gonic/gin"
)

// @Summary 		Create Author
// @Description 	Create page
// @Tags 			Author
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.AuthorReq    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/author/create [post]
func (h *Handler) CreateAuthor(ctx *gin.Context) {
	arr := &pb.AuthorReq{}
	err := ctx.BindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.Author.CreateAuthor(ctx, arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Update Author
// @Description 	Update page
// @Tags 			Author
// @Accept  		json
// @Produce  		json
// @Param     		id 		path    string   true    "Author ID"
// @Param   		Update  body    pb.AuthorRes  true    "Update"
// @Success 		200   {string}  string   "Update Successful"
// @Failure 		401   {string}  string   "Error while updated"
// @Router			 /author/update/{id} [put]
func (h *Handler) UpdateAuthor(ctx *gin.Context) {
	arr := pb.AuthorRes{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	arr.Id = ctx.Param("id")

	_, err = h.Author.UpdateAuthor(ctx, &arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Delete Author
// @Description 	Delete page
// @Tags 			Author
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true    "Author ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/author/delete/{id} [delete]
func (h *Handler) DeleteAuthor(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Author.DeleteAuthor(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll Author
// @Description 	GetAll page
// @Tags			Author
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.FilterAuthor true "Query parameter"
// @Success 		200  {object}  pb.AllAuthors     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAll"
// @Router 			/author/getall [get]
func (h *Handler) GetAllAuthors(ctx *gin.Context) {
	author := &pb.FilterAuthor{}
	author.Name = ctx.Query("name")

	res, err := h.Author.GetAllAuthors(ctx, author)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		GetById Author
// @Description 	GetById page
// @Tags 			Author
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string      true   "Author ID"
// @Success 		200  {object}  pb.AuthorRes "GetById Successful"
// @Failure 		401  {string}  string     "Error while GetById"
// @Router 			/author/get/{id} [get]
func (h *Handler) GetByIdAuthor(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Author.GetByIdAuthor(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		GetAll Author
// @Description 	GetAll page
// @Tags			Author
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string      true   "Author ID"
// @Success 		200  {object}  pb.AllBooks     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAll"
// @Router 			/author/{id}/books [get]
func (h *Handler) GetAuthorBooks(ctx *gin.Context) {
	books, err := h.Book.GetAllBooks(ctx, &pb.FilterBook{AuthorId: ctx.Param("id")})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, books)
}