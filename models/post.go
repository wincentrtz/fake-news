package models

type Post struct {
	Id          int    `json:"id"`
	Parent      int    `json:"post_parent_id"`
	Title       string `json:"post_title"`
	Description string `json:"post_description"`
}
