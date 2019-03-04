package routers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/goalong/bingo/utils"
	"github.com/goalong/bingo/models"
	myError "github.com/goalong/bingo/err"
)

// 获取post列表
func GetPosts(c *gin.Context) {
	filter := make(map[string]interface{})
	ret := make(map[string]interface{})

	page, pageSize := utils.GetPageInfo(c)
	posts := models.GetPosts(page, pageSize, filter)
	ret["items"] = posts
	ret["total"] = 10 // todo
	code := myError.SUCCESS


	c.JSON(
		200,
		gin.H{
			"code": code,
			"msg": "", //todo
			"data": ret,
		})

}

//新建post
func AddPost(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	statusCode := 200
	code := myError.SUCCESS
	if err != nil {
		code = myError.INVALID_PARAMS
		statusCode = 400
	}
	err = models.CreatePost(post)
	if err != nil {
		code = myError.INVALID_PARAMS
		statusCode = 400
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": "", //todo
	})
}

func EditPost(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	id := com.StrTo(c.Param("id")).MustInt()
	statusCode := 200
	code := myError.SUCCESS

	if err != nil {
		code = myError.INVALID_PARAMS
		statusCode = 400
	}
	models.EditPost(id, post)
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": "", //todo
	})

}

