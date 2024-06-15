package handler

import (
	pb "github.com/Salikhov079/library-api/genprotos"

	"github.com/gin-gonic/gin"
)

// @Summary 		Create Genre
// @Description 	Create page
// @Tags 			Genre
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Genre    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/genre/create [post]
func (h *Handler) CreateGenre(ctx *gin.Context) {
	arr := &pb.Genre{}
	err := ctx.BindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.Genre.CreateGenre(ctx, arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Update Genre
// @Description 	Update page
// @Tags 			Genre
// @Accept  		json
// @Produce  		json
// @Param     		id 		path    string   true    "Genre ID"
// @Param   		Update  body    pb.Genre  true    "Update"
// @Success 		200   {string}  string   "Update Successful"
// @Failure 		401   {string}  string   "Error while updated"
// @Router			 /genre/update/{id} [put]
func (h *Handler) UpdateGenre(ctx *gin.Context) {
	arr := pb.Genre{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	arr.Id = ctx.Param("id")

	_, err = h.Genre.UpdateGenre(ctx, &arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Delete Genre
// @Description 	Delete page
// @Tags 			Genre
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true    "Genre ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/genre/delete/{id} [delete]
func (h *Handler) DeleteGenre(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Genre.DeleteGenre(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll Genre
// @Description 	GetAll page
// @Tags			Genre
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Genre true "Query parameter"
// @Success 		200  {object}  pb.AllGenres     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAlld"
// @Router 			/genre/getall [get]
func (h *Handler) GetAllGenres(ctx *gin.Context) {
	genre := &pb.Genre{}
	genre.Name = ctx.Param("name")

	res, err := h.Genre.GetAllGenres(ctx, genre)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		GetById Genre
// @Description 	GetById page
// @Tags 			Genre
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string      true   "Genre ID"
// @Success 		200  {object}  pb.Genre "GetById Successful"
// @Failure 		401  {string}  string     "Error while GetByIdd"
// @Router 			/genre/get/{id} [get]
func (h *Handler) GetByIdGenre(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Genre.GetByIdGenre(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}
