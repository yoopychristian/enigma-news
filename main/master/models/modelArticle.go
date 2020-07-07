package models

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	ArticleText string `json:"article"`
}
