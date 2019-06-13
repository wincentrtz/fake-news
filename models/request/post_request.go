package request

type PostRequest struct {
	Parent      int    `json:"post_parent_id"`
	Title       string `json:"post_title"`
	Description string `json:"post_description"`
}
