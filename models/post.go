package models

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"post_title"`
	Content string `json:"post_description"`
}
