package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoafuzzz/blog-go/entity"
	"github.com/gogoafuzzz/blog-go/service"
)

type ArticleController interface {
	FindAll() []entity.Article
	Save(ctx *gin.Context) entity.Article
}

type controller struct {
	service service.ArticleService
}

func New(service service.ArticleService) ArticleController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Article {
	var article entity.Article
	ctx.BindJSON(&article)
	return c.service.Save(article)
}

func (c *controller) FindAll() []entity.Article {
	return c.service.FindAll()
}
