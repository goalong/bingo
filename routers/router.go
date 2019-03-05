package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/goalong/bingo/utils"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug") // todo
	r.POST("/login", Login)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(utils.CheckToken())

	apiv1.GET("/posts", GetPosts)
	apiv1.GET("/post/:id", GetPost)
	apiv1.POST("post", AddPost)
	apiv1.PUT("/post/:id", EditPost)
	apiv1.DELETE("/post/:id", DeletePost)
	apiv1.POST("/register", UserRegister)


	return r
}
