package main

import (
	"net/http"
	"time"

	"github.com/MRdyRy/go-crud-food/config"
	"github.com/MRdyRy/go-crud-food/controller"
	"github.com/MRdyRy/go-crud-food/helper"
	"github.com/MRdyRy/go-crud-food/model"
	"github.com/MRdyRy/go-crud-food/repository"
	"github.com/MRdyRy/go-crud-food/router"
	"github.com/MRdyRy/go-crud-food/service"

	"github.com/go-playground/validator/v10"
)

const (
	host         = "localhost:"
	port         = "8080"
	readTimeOut  = 10 * time.Second
	writeTimeOut = 10 * time.Second
	tableName    = "foods"
)

func main() {

	//DB
	db := config.DBConn()
	validate := validator.New()

	db.Table(tableName).AutoMigrate(&model.Food{})

	//init repository
	foodRepository := repository.NewFoodRepositoryImpl(db)

	//init service
	foodService := service.NewFoodServiceImpl(foodRepository, validate)

	//init controller
	foodController := controller.NewFoodController(foodService)

	//init routers
	routes := router.NewRouter(foodController)

	server := &http.Server{
		Addr:           host + port,
		Handler:        routes,
		ReadTimeout:    readTimeOut,
		WriteTimeout:   writeTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
