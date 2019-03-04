package routers

import (
	"github.com/gin-gonic/gin"
	"../utils"
	"../models"
	"../err"
)

func GetPosts(c *gin.Context) {
	filter := make(map[string]interface{})
	ret := make(map[string]interface{})

	page, pageSize := utils.GetPageInfo(c)
	posts := models.GetPosts(page, pageSize, filter)
	ret["items"] = posts
	ret["total"] = 10 // todo
	code := err.SUCCESS


	c.JSON(
		200,
		gin.H{
			"code": code,
			"msg": "", //todo
			"data": ret,
		})

}

func AddPost(c *gin.Context) {

}

