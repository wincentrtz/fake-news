package models

type PostQueue struct {
	Id       int `json:"id"`
	PostId   int `json:"post_id"`
	Progress int `json:"progress"`
}
