package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回代码
// 请求相关 0-99
// 用户相关 100-199
// 查询相关 200-299
const (
	Error        = -1  // 失败
	Success      = 0   // 成功
	Unauthorized = 100 // 未认证
	NotFound     = 200 // 未查询到相关数据
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
