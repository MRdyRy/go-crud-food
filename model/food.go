package model

type Food struct {
	Id    int     `gorm:"type:int;primary_key"`
	Name  string  `gorm:"type:varchar(255)"`
	Price float32 `gorm:"type:decimal(10,2)"`
}
