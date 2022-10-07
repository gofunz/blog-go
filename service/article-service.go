package service

import "github.com/gogoafuzzz/blog-go/entity"

type ArticleService interface {
	Save(entity.Article) entity.Article
	FindAll() []entity.Article
}

type articleService struct {
	articles []entity.Article
}

// new article service instance
func New() ArticleService {
	return &articleService{}
}

// save an article
func (service *articleService) Save(article entity.Article) entity.Article {
	service.articles = append(service.articles, article)
	return article
}

// find all articles
func (service *articleService) FindAll() []entity.Article {
	return service.articles
}
