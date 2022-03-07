package models

type Book struct {
	Model
	UserID   uint   `json:"-" gorm:"comment:外键"`
	Title    string `json:"title" gorm:"not null;comment:书名" binding:"required"`
	IpfsKey  string `json:"ipfs_key" gorm:"not null;comment:IPFS key" binding:"required"`
	FileType string `json:"file_type" gorm:"not null;comment:文件类型" binding:"required"`
}
