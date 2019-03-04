package utils

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// 获取页码及每页数量
func GetPageInfo(c *gin.Context) (page int, pageSize int) {
	page, _ = com.StrTo(c.Query("page")).Int()
	pageSize, _ = com.StrTo(c.Query("pageSize")).Int()
	return page, pageSize
}