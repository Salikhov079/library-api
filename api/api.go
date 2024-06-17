package api

import (
	"github.com/Salikhov079/library-api/api/handler"
	_ "github.com/Salikhov079/library-api/docs"
	"github.com/gin-gonic/gin"


	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Library service
// @version 1.0
// @description Library service
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	book := r.Group("/book")
	book.POST("/create", h.CreateBook)
	book.PUT("/update/:id", h.UpdateBook)
	book.DELETE("/delete/:id", h.DeleteBook)
	book.GET("/getall", h.GetAllBooks)
	book.GET("/get/:id", h.GetByIdBook)
	book.GET("/overdue", h.OverdueBooks)

	author := r.Group("/author")
	author.POST("/create", h.CreateAuthor)
	author.PUT("/update/:id", h.UpdateAuthor)
	author.DELETE("/delete/:id", h.DeleteAuthor)
	author.GET("/getall", h.GetAllAuthors)
	author.GET("/get/:id", h.GetByIdAuthor)
	author.GET("/:id/books", h.GetAuthorBooks)

	genre := r.Group("/genre")
	genre.POST("/create", h.CreateGenre)
	genre.PUT("/update/:id", h.UpdateGenre)
	genre.DELETE("/delete/:id", h.DeleteGenre)
	genre.GET("/getall", h.GetAllGenres)
	genre.GET("/get/:id", h.GetByIdGenre)
	genre.GET("/:id/books", h.GetGenreBooks)

	borrower := r.Group("/borrower")
	borrower.POST("/create", h.CreateBorrower)
	borrower.PUT("/update/:id", h.UpdateBorrower)
	borrower.DELETE("/delete/:id", h.DeleteBorrower)
	borrower.GET("/getall", h.GetAllBorrowers)
	borrower.GET("/get/:id", h.GetByIdBorrower)

	return r
}
