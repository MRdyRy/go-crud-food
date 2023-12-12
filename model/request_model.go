package model

type Request struct {
	Name  string  `validate:"required,min=1,max=200" json:"name"`
	Price float32 `json:"price"`
}
