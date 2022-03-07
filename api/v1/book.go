package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ipfs-library-server/common/resp"
	"ipfs-library-server/common/utils"
	"ipfs-library-server/middleware"
	"ipfs-library-server/models"
	"ipfs-library-server/service"
	"ipfs-library-server/structs"
	"strconv"
)

// AddBook 添加书籍
func AddBook(c *gin.Context) {
	book := models.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		resp.Result(resp.Error, utils.Translate(err), nil, c)
		return
	}
	book.UserID = middleware.GetUserID(c)
	if err := service.AddBook(&book); err != nil {
		resp.Result(resp.Error, utils.Translate(err), nil, c)
		return
	}
	resp.Result(resp.Success, fmt.Sprintf("Add books《%s》success", book.Title), nil, c)
}

// SearchBooks 查询书籍
func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	totalRows, userList, err := service.SearchBooks(title, pageSize, pageNum)
	if err != nil {
		resp.Result(resp.Error, "Failed to query books", nil, c)
		return
	}
	if totalRows == 0 {
		resp.Result(resp.NotFound, "Book not found", nil, c)
		return
	}
	data := structs.Pagination{
		TotalRows: totalRows,
		Data:      userList,
	}
	resp.Result(resp.Success, "Book query success", data, c)
}

// QueryBooksList 查询某个用户添加的所有书籍
func QueryBooksList(c *gin.Context) {

}

// DeleteBook 删除某本书
func DeleteBook(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteBook(bookId); err != nil {
		resp.Result(resp.NotFound, "Book delete failed", nil, c)
		return
	}
	resp.Result(resp.Success, "Book delete success", nil, c)
}
