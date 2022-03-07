package models

type Region struct {
	Model
	Name string `json:"name" gorm:"unique;not null;comment:区域名" binding:"required"`
}
