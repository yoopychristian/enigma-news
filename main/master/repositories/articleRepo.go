package repositories

import "enigma-news/main/master/models"

type ArticleRepository interface {
	GetAllArticle() ([]*models.Article, error)
	GetArticleID(Id string) (*models.Article, error)
	CreateArticle(article models.Article) error
	UpdateArticle(article models.Article) error
	DeleteArticle(Id string) error
}
