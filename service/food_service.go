package service

import "github.com/MRdyRy/go-crud-food/model"

type FoodService interface {
	Create(f model.Request)
	Update(f model.UpdateReq)
	Delete(fId int)
	FindById(fId int) model.Response
	FindAll() []model.Response
}
