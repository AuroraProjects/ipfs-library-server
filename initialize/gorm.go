package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ipfs-library-server/global"
	"ipfs-library-server/models"
)

func Gorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		global.Zap.Panicf("Failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(
		models.User{},
		models.Book{},
		models.Region{},
		models.Gateway{},
	); err != nil {
		global.Zap.Panicf("Auto migrate failed: %v", err)
	}

	return db
}
