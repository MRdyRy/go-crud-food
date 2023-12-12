package controller

import (
	"go-crud-food/helper"
	"go-crud-food/model"
	"go-crud-food/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	foodService service.FoodService
}

func NewFoodController(service service.FoodService) *FoodController {
	return &FoodController{foodService: service}
}

func (controller *FoodController) Create(ctx *gin.Context) {
	createFoodReq := model.Request{}
	err := ctx.ShouldBindJSON(&createFoodReq)
	helper.ErrorPanic(err)

	controller.foodService.Create(createFoodReq)

	webResponse := model.ResponseWeb{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FoodController) Update(ctx *gin.Context) {
	updateReq := model.UpdateReq{}
	err := ctx.ShouldBindJSON(&updateReq)
	helper.ErrorPanic(err)

	foodId := ctx.Param("foodId")

	updateReq.Id = helper.GetIdFromString(foodId)

	controller.foodService.Update(updateReq)

	webResponse := model.ResponseWeb{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *FoodController) Delete(ctx *gin.Context) {
	foodId := ctx.Param("foodId")
	id := helper.GetIdFromString(foodId)

	controller.foodService.Delete(id)
	webResponse := model.ResponseWeb{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FoodController) FindById(ctx *gin.Context) {
	foodId := ctx.Param("foodId")
	id := helper.GetIdFromString(foodId)
	foodResponse := controller.foodService.FindById(id)

	webResponse := model.ResponseWeb{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   foodResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FoodController) FindAll(ctx *gin.Context) {
	foodResponse := controller.foodService.FindAll()

	webResponse := model.ResponseWeb{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   foodResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)

}
