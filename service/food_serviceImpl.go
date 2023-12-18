package service

import (
	"github.com/MRdyRy/go-crud-food/helper"
	"github.com/MRdyRy/go-crud-food/model"
	"github.com/MRdyRy/go-crud-food/repository"

	"github.com/go-playground/validator/v10"
)

type FoodServiceImpl struct {
	FoodRepository repository.FoodRepository
	Validate       *validator.Validate
}

// Create implements FoodService.
func (t *FoodServiceImpl) Create(f model.Request) {
	err := t.Validate.Struct(f)
	helper.ErrorPanic(err)
	foodModel := model.Food{
		Name:  f.Name,
		Price: f.Price,
	}
	t.FoodRepository.Save(foodModel)
}

// Delete implements FoodService.
func (t *FoodServiceImpl) Delete(fId int) {
	t.FoodRepository.Delete(fId)
}

// FindAll implements FoodService.
func (t *FoodServiceImpl) FindAll() []model.Response {
	results := t.FoodRepository.FindAll()

	var foods []model.Response

	for _, value := range results {
		food := model.Response{
			Id:    value.Id,
			Name:  value.Name,
			Price: value.Price,
		}
		foods = append(foods, food)
	}
	return foods
}

// FindById implements FoodService.
func (t *FoodServiceImpl) FindById(fId int) model.Response {
	foodData, err := t.FoodRepository.FindById(fId)
	helper.ErrorPanic(err)
	foodResponse := model.Response{
		Id:    foodData.Id,
		Name:  foodData.Name,
		Price: foodData.Price,
	}
	return foodResponse
}

// Update implements FoodService.
func (t *FoodServiceImpl) Update(f model.UpdateReq) {
	foodData, err := t.FoodRepository.FindById(f.Id)
	helper.ErrorPanic(err)
	foodData.Name = f.Name
	foodData.Price = f.Price
	t.FoodRepository.Update(foodData)

}

func NewFoodServiceImpl(foodRepository repository.FoodRepository, validate *validator.Validate) FoodService {
	return &FoodServiceImpl{
		FoodRepository: foodRepository,
		Validate:       validate,
	}
}
