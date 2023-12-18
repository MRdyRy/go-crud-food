package repository

import "github.com/MRdyRy/go-crud-food/model"

type FoodRepository interface {
	Save(f model.Food)
	Delete(fId int)
	Update(f model.Food)
	FindById(fId int) (f model.Food, err error)
	FindAll() []model.Food
}
