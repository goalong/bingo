package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug") // todo
	r.GET("/posts", GetPosts)
	r.GET("/post/:id", GetPost)
	r.POST("post", AddPost)
	r.PUT("/post/:id", EditPost)
	r.DELETE("/post/:id", DeletePost)

	return r
}
