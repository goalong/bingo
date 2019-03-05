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
			"msg": myError.GetMsg(code),
			"data": ret,
		})

}

func GetPost(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	post := models.GetPost(id)
	code := myError.SUCCESS
	if post.Id < 1 {
		code = myError.PostNotFound
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg": myError.GetMsg(code),
		"data": post,
	})


}

//新建post
func AddPost(c *gin.Context) {
	var post models.Post
	statusCode := 200
	code := myError.SUCCESS
	user, ok  := c.Get("user")
	if !ok {
		code = myError.CONTEXT_GET_USER_ERROR

	} else {
		err := c.BindJSON(&post) // json的解析
		if err != nil {
			code = myError.INVALID_PARAMS
			statusCode = 400
		}
		u, _ := user.(models.User) // 接口的类型断言，这里将接口类型转换成了User struct
		post.UserId = u.ID
		err = models.CreatePost(post)
		if err != nil {
			code = myError.INVALID_PARAMS
			statusCode = 400
		}

	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": myError.GetMsg(code),
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
	user, _  := c.Get("user")
	u, _ := user.(models.User)
	p := models.GetPost(id)
	// 检查是否修改自己的文章
	if u.ID != p.UserId {
		code = myError.PERMISSION_DENIED
		statusCode = 403
	} else {
		update := make(map[string]interface{})
		//只有title和description两个字段允许修改
		update["title"] = post.Title
		update["description"] = post.Description
		models.EditPost(id, update)
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": myError.GetMsg(code),
	})

}


func DeletePost(c *gin.Context) {
	code := myError.SUCCESS
	id := com.StrTo(c.Param("id")).MustInt()
	isExist := models.IsPostExist(id)
	statusCode := 200
	if !isExist {
		code = myError.PostNotFound
	} else {
		user, _  := c.Get("user")
		u, _ := user.(models.User)
		p := models.GetPost(id)
		// 检查操作的文章是否属于该用户
		if u.ID != p.UserId {
			code = myError.PERMISSION_DENIED
			statusCode = 403
		} else {
			models.DeletePost(id)
		}
	}
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": myError.GetMsg(code),
	})
}

