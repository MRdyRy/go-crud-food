package helper

import (
	"strconv"
)

func GetIdFromString(id string) int {
	fId, err := strconv.Atoi(id)
	ErrorPanic(err)
	foodId := fId
	return foodId
}
