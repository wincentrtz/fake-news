package models

type PostQueue struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	PostTitle string `json:"post_title"`
	Progress  int    `json:"progress"`
}
