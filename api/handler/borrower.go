package handler

import (
	"fmt"

	pb "github.com/Salikhov079/library-api/genprotos"

	"github.com/gin-gonic/gin"
)

// @Summary 		Create Borrower
// @Description 	Create page
// @Tags 			Borrower
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.BorrowerReq   true   "Create"
// @Success 		200   {string}      string      "Create Successful"
// @Failure 		401   {string}      string      "Error while Created"
// @Router 			/borrower/create [post]
func (h *Handler) CreateBorrower(ctx *gin.Context) {
	arr := &pb.BorrowerReq{}
	err := ctx.BindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.Borrower.CreateBorrower(ctx, arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Update Borrower
// @Description 	Update page
// @Tags 			Borrower
// @Accept  		json
// @Produce  		json
// @Param     		id 		path    string   true    "Borrower ID"
// @Param   		Update  body    pb.BorrowerRes  true    "Update"
// @Success 		200   {string}  string   "Update Successful"
// @Failure 		401   {string}  string   "Error while updated"
// @Router			 /borrower/update/{id} [put]
func (h *Handler) UpdateBorrower(ctx *gin.Context) {
	arr := pb.BorrowerRes{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	arr.Id = ctx.Param("id")

	_, err = h.Borrower.UpdateBorrower(ctx, &arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		Delete Borrower
// @Description 	Delete page
// @Tags 			Borrower
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true    "Borrower ID"
// @Success 		200 {string}  string  "Delete Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/borrower/delete/{id} [delete]
func (h *Handler) DeleteBorrower(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Borrower.DeleteBorrower(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll Borrower
// @Description 	GetAll page
// @Tags			Borrower
// @Accept  		json
// @Produce  		json
// @Success 		200  {object}  pb.AllBorrowers     "GetAll Successful"
// @Failure 		401  {string}  string          "Error while GetAlld"
// @Router 			/borrower/getall [get]
func (h *Handler) GetAllBorrowers(ctx *gin.Context) {
	res, err := h.Borrower.GetAllBorrowers(ctx, &pb.BorrowerReq{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resU, err := h.User.GetAllUsers(ctx, &pb.UserReq{})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	
	check(res, resU)

	ctx.JSON(200, res)
}

func check(res *pb.AllBorrowers, resU *pb.AllUsers) {
	for i := 0; i < len(res.Borrowers); i++ {
		for j := 0; j < len(resU.Users); j++ {
			if res.Borrowers[i].User.Id == resU.Users[j].Id {
				res.Borrowers[i].User = resU.Users[j]
			}
		}
	}
}

// @Summary 		GetById Borrower
// @Description 	GetById page
// @Tags 			Borrower
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string      true   "Borrower ID"
// @Success 		200  {object}  pb.BorrowerRes "GetById Successful"
// @Failure 		401  {string}  string     "Error while GetByIdd"
// @Router 			/borrower/get/{id} [get]
func (h *Handler) GetByIdBorrower(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Borrower.GetByIdBorrower(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	res.User, err = h.User.GetByIdUser(ctx, &pb.ById{Id: res.User.Id})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	fmt.Println(res)
	ctx.JSON(200, res)
}
