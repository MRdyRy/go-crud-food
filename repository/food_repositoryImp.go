package repository

import (
	"errors"

	"github.com/MRdyRy/go-crud-food/helper"
	"github.com/MRdyRy/go-crud-food/model"

	"gorm.io/gorm"
)

type FoodRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements FoodRepository.
func (t *FoodRepositoryImpl) Delete(fId int) {
	var f model.Food
	result := t.Db.Where("id = ?", fId).Delete(&f)
	helper.ErrorPanic(result.Error)
}

// FindAll implements FoodRepository.
func (t *FoodRepositoryImpl) FindAll() []model.Food {
	var foods []model.Food
	results := t.Db.Find(&foods)
	helper.ErrorPanic(results.Error)
	return foods
}

// FindById implements FoodRepository.
func (t *FoodRepositoryImpl) FindById(fId int) (model.Food, error) {
	var f model.Food
	result := t.Db.Find(&f, fId)
	if result != nil {
		return f, nil
	} else {
		return f, errors.New("food is not found")
	}
}

// Save implements FoodRepository.
func (t *FoodRepositoryImpl) Save(f model.Food) {
	result := t.Db.Create(&f)
	helper.ErrorPanic(result.Error)
}

// Update implements FoodRepository.
func (t *FoodRepositoryImpl) Update(f model.Food) {
	var updateF = model.Request{
		Name:  f.Name,
		Price: f.Price,
	}

	result := t.Db.Model(&f).Updates(updateF)
	helper.ErrorPanic(result.Error)
}

func NewFoodRepositoryImpl(Db *gorm.DB) FoodRepository {
	return &FoodRepositoryImpl{Db: Db}
}
