package controllers

import (
	"encoding/json"
	"enigma-news/main/master/handler"
	"enigma-news/main/master/middleware"
	"enigma-news/main/master/models"
	"enigma-news/main/master/response"
	"enigma-news/main/master/usecases"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	articleUseCase usecases.ArticleUseCase
}

func ArticleController(r *mux.Router, service usecases.ArticleUseCase) {
	r.Use(middleware.ActivityLogMiddleware)
	ArticleHandler := ArticleHandler{service}

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", handler.NewAuthenticationHandler().Handler).Methods(http.MethodPost)
	// auth.HandleFunc("/tokenValidation", handler.NewTokenValidationHandler().Handler).Methods(http.MethodGet)

	articles := r.PathPrefix("/articles").Subrouter()
	articles.Use(middleware.TokenValidationMiddleware)
	articles.HandleFunc("", ArticleHandler.ListArticles).Methods(http.MethodGet)

	article := r.PathPrefix("/article").Subrouter()
	article.Use(middleware.TokenValidationMiddleware)
	article.HandleFunc("/{id}", ArticleHandler.GetArticleID).Methods(http.MethodGet)
	article.HandleFunc("/add", ArticleHandler.CreateDataArticles).Methods(http.MethodPost)
	article.HandleFunc("/update/{id}", ArticleHandler.UpdateDataArticles).Methods(http.MethodPut)
	article.HandleFunc("/delete{id}", ArticleHandler.DeleteDataArticles).Methods(http.MethodDelete)

}

func (s ArticleHandler) ListArticles(w http.ResponseWriter, r *http.Request) {
	article, err := s.articleUseCase.GetAllArticle()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Articles"
	pesan.Data = article

	byteOfArticles, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfArticles)
	fmt.Println("EndPoint Hit : GetArticlePage")
}

func (s ArticleHandler) GetArticleID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	article, err := s.articleUseCase.GetArticleID(vars["Id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Get Database Article by ID"
	pesan.Data = article

	byteOfArticles, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArticles)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: GetArticleHByIDPage")
}

func (s ArticleHandler) CreateDataArticles(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article) // json ke struct
	s.articleUseCase.CreateArticle(article)
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Create Data for Article Database"
	pesan.Data = "Success"

	byteOfArticles, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArticles)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: CreateArticlePage")
}

func (s ArticleHandler) UpdateDataArticles(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article) // json ke struct
	s.articleUseCase.UpdateArticle(article)

	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Update Data for Article Database"
	pesan.Data = "Success"

	byteOfArticles, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArticles)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: UpdateArticlePage")
}

func (s ArticleHandler) DeleteDataArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.articleUseCase.DeleteArticle(vars["Id"])
	var pesan response.Response
	pesan.Status = http.StatusOK
	pesan.Messages = "Delete Database Article"
	pesan.Data = "Success"

	byteOfArticles, err := json.Marshal(pesan)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArticles)
	w.WriteHeader(http.StatusAlreadyReported)
	fmt.Println("Endpoint Hit: DeleteArticlePage")
}
