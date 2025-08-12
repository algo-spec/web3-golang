package routes

import (
	"task04/app/handlers"
	"task04/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户模块路由
	user := r.Group("/user")
	{
		user.POST("/register", handlers.Register)
		user.POST("/login", handlers.Login)
	}

	// 文章模块路由
	post := r.Group("/post")
	{
		post.GET("", handlers.GetPosts)
		post.GET("/:id", handlers.GetPost)

		autoPost := post.Group("")
		autoPost.Use(middleware.JWTAuth())
		{
			autoPost.POST("/create", handlers.CreatePost)
			autoPost.PUT("/:id", handlers.UpdatePost)
			autoPost.DELETE("/:id", handlers.DeletePost)
		}
	}

	// 评论模块路由
	comment := r.Group("/comment")
	{
		comment.GET("/:post_id", handlers.GetComment)
		comment.Use(middleware.JWTAuth())
		comment.POST("/create", handlers.CreateComment)
	}
}
