package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gogoafuzzz/blog-go/entity"
	"github.com/gogoafuzzz/blog-go/service"
	"github.com/gogoafuzzz/blog-go/validations"
)

type ArticleController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.ArticleService
}

var validator *baseValidator.Validate

func registerValidations() {
	validator = baseValidator.New()
	validator.RegisterValidation("is-cool", validations.ValidateTitleIsCool)
}

func New(service service.ArticleService) ArticleController {
	registerValidations()
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) {
	var article entity.Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		var errors = strings.Split(string(err.Error()), "\n")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	// 透過 validator 驗證自定義的 validation "is-cool"
	err = validator.Struct(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.service.Save(article)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": article,
	})
}

func (c *controller) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.FindAll())
}

func (c *controller) ShowAll(ctx *gin.Context) {
	articles := c.service.FindAll()
	data := gin.H{
		"title":    "Article Page",
		"articles": articles,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
