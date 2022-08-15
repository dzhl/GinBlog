package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		// router.GET("hello", func(ctx *gin.Context) {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })
		//用户模块路由
		Auth.POST("user/add", v1.AddUser)
		Auth.GET("users", v1.GetUsers)
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块路由
		Auth.POST("category/add", v1.AddCategory)

		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块路由
		Auth.POST("article/add", v1.AddArticle)

		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)
	}
	router := r.Group("api/v1")
	{

		//用户模块路由
		//Auth.POST("user/add", v1.AddUser)
		//Auth.GET("users", v1.GetUsers)
		//Auth.PUT("user/:id", v1.EditUser)
		//Auth.DELETE("user/:id", v1.DeleteUser)
		router.POST("user/login", v1.Login)
		//分类模块路由
		//Auth.POST("category/add", v1.AddCategory)
		router.GET("categorys", v1.GetCategorys)
		//Auth.PUT("category/:id", v1.EditCategory)
		//Auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块路由
		//router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticles)
		router.GET("article/:id", v1.GetArticleInfo)
		router.GET("articles/:id", v1.GetCatArticles)
		//router.PUT("article/:id", v1.EditArticle)
		//router.DELETE("article/:id", v1.DeleteArticle)

	}
	r.Run(utils.HttpPort)
}
