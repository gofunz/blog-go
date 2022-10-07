package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogoafuzzz/blog-go/controller"
	"github.com/gogoafuzzz/blog-go/service"
)

var (
	articleService    service.ArticleService       = service.New()
	articleController controller.ArticleController = controller.New(articleService)
)

func main() {

	router := gin.Default()

	articles := router.Group("articles")
	{
		articles.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, articleController.FindAll())
		})

		articles.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, articleController.Save(ctx))
		})
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8000")
}
