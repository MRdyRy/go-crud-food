package model

type UpdateReq struct {
	Id    int     `validate:"required"`
	Name  string  `validate:"required,max=200,min=1" json:"name"`
	Price float32 `validate:"required" json:"price"`
}
