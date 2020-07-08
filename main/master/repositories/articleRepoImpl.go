package repositories

import (
	"database/sql"
	"enigma-news/main/master/constantaQuery"
	"enigma-news/main/master/models"
	"log"
)

type ArticleRepoImpl struct {
	db *sql.DB
}

func (s ArticleRepoImpl) GetAllArticle() ([]*models.Article, error) {
	dataArticle := []*models.Article{}
	query := constantaQuery.GETALLARTICLE
	data, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for data.Next() {
		article := models.Article{}
		var err = data.Scan(&article.Id, &article.Title, &article.ArticleText)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		dataArticle = append(dataArticle, &article)
	}
	return dataArticle, nil
}

func (s ArticleRepoImpl) GetArticleID(Id string) (*models.Article, error) {
	article := new(models.Article)
	query := constantaQuery.GETARTICLEBYID
	if err := s.db.QueryRow(query, Id).Scan(&article.Id, &article.Title, &article.ArticleText); err != nil {
		log.Println(err)
		return nil, err
	}
	return article, nil
}

func (s ArticleRepoImpl) CreateArticle(article models.Article) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	query := constantaQuery.GETADDARTICLE

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(article.Id, article.Title, article.ArticleText); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func (s ArticleRepoImpl) UpdateArticle(article models.Article) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETUPDATEARTICLE

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(article.Title, article.ArticleText, article.Id); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func (s ArticleRepoImpl) DeleteArticle(Id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETUPDATEARTICLE

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(Id); err != nil {
		tx.Rollback()
		log.Println(err)
		log.Fatalf("%v", err)
		return err
	}
	return nil
}

func InitArticleRepoImpl(db *sql.DB) ArticleRepository {
	return &ArticleRepoImpl{db}
}
