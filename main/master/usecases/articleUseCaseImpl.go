package usecases

import (
	"enigma-news/main/master/models"
	"enigma-news/main/master/repositories"
	"enigma-news/main/master/utils"
)

type ArticleUsecaseImpl struct {
	articleRepo repositories.ArticleRepository
}

func (s ArticleUsecaseImpl) GetAllArticle() ([]*models.Article, error) {
	article, err := s.articleRepo.GetAllArticle()
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s ArticleUsecaseImpl) GetArticleID(Id string) (*models.Article, error) {
	article, err := s.articleRepo.GetArticleID(Id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s ArticleUsecaseImpl) CreateArticle(article models.Article) error {
	err := utils.ValidateInputNotNil(article.Title, article.ArticleText)
	if err != nil {
		return err
	}
	err = s.articleRepo.CreateArticle(article)
	if err != nil {
		return err
	}
	return nil
}

func (s ArticleUsecaseImpl) UpdateArticle(article models.Article) error {
	err := utils.ValidateInputNotNil(article.Title, article.ArticleText)
	if err != nil {
		return err
	}
	err = s.articleRepo.UpdateArticle(article)
	if err != nil {
		return err
	}
	return nil
}
func (s ArticleUsecaseImpl) DeleteArticle(Id string) error {
	err := s.articleRepo.DeleteArticle(Id)
	if err != nil {
		return err
	}
	return nil
}

func InitArticleUseCase(articleRepo repositories.ArticleRepository) ArticleUseCase {
	return &ArticleUsecaseImpl{articleRepo}
}
