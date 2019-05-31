package models

type Post struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
