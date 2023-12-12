package router

import (
	"go-crud-food/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(foodController *controller.FoodController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World !")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	foodRouter := router.Group("/food")
	foodRouter.GET("/all", foodController.FindAll)
	foodRouter.GET("/get/:foodId", foodController.FindById)
	foodRouter.POST("/create", foodController.Create)
	foodRouter.PATCH("/patch/:foodId", foodController.Update)
	foodRouter.DELETE("/delete/:foodId", foodController.Delete)
	return service
}
