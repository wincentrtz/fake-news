package models

type PostStatus struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	PostTitle string `json:"post_title"`
	Status    int    `json:"status"`
}
