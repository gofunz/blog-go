package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogoafuzzz/blog-go/controller"
	"github.com/gogoafuzzz/blog-go/middleware"
	"github.com/gogoafuzzz/blog-go/service"
)

var (
	articleService    service.ArticleService       = service.New()
	articleController controller.ArticleController = controller.New(articleService)
)

var fp *os.File

// 將 gin 請求記錄寫到日誌中
func setupLogOutput() {
	fp, _ = os.Create("http.log")
	gin.DefaultWriter = io.MultiWriter(fp, os.Stdout)
}

func main() {

	setupLogOutput()
	defer fp.Close()

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		articles := apiRoutes.Group("articles")
		{
			articles.GET("/", articleController.FindAll)
			articles.POST("/", articleController.Save)
		}
	}

	webRoutes := router.Group("")
	{
		webRoutes.GET("/", articleController.ShowAll)
	}

	router.Run(":8000")
}
