package models

type Gateway struct {
	Model
	UserID        uint   `json:"-" gorm:"comment:外键"`
	RegionId      int    `json:"region_id" gorm:"not null;comment:区域ID" binding:"required"`
	GatewayDomain string `json:"gateway_domain" gorm:"unique;not null;comment:网关地址" binding:"required"`
}
