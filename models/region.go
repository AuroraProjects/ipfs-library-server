package models

type Region struct {
	Model
	Name string `json:"name" gorm:"unique;not null;comment:εΊεε" binding:"required"`
}
