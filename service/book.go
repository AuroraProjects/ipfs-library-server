package service

import (
	"gorm.io/gorm"
	"ipfs-library-server/global"
	"ipfs-library-server/models"
)

// AddBook 添加书籍到数据库
func AddBook(book *models.Book) error {
	if err := global.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

// SearchBooks 搜索书籍
func SearchBooks(title string, pageSize int, pageNum int) (int64, []models.Book, error) {
	var totalRows int64
	var books []models.Book
	condition := []string{"id", "title", "ipfs_key", "file_type", "created_at", "updated_at"}
	err := global.DB.Select(condition).Where("title = ?", title).Find(&books).
		Count(&totalRows).Limit(pageNum).Offset(pageNum * (pageSize - 1)).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, nil, err
	}
	return totalRows, books, nil
}

// QueryBooksList 查询用户下的所有书籍
func QueryBooksList(userID uint, pageSize int, pageNum int) (int64, []models.Book, error) {
	var books []models.Book
	var totalRows int64
	err := global.DB.Where("user_id = ?", userID).Find(&books).
		Count(&totalRows).Limit(pageNum).Offset(pageNum * (pageSize - 1)).Error
	if err != nil {
		return 0, nil, err
	}
	return 0, books, nil
}

// DeleteBook 删除书籍
func DeleteBook(bookID int) error {
	var book = models.Book{}
	if err := global.DB.Delete(&book, bookID).Error; err != nil {
		return err
	}

	return nil
}
