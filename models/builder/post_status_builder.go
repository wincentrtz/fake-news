package builder

import "github.com/wincentrtz/fake-news/models"

type postStatusBuilder struct {
	id        int
	postId    int
	postTitle string
	status    int
}

// PostStatusBuilder builder interface
type PostStatusBuilder interface {
	Id(int) PostStatusBuilder
	PostId(int) PostStatusBuilder
	PostTitle(string) PostStatusBuilder
	Status(int) PostStatusBuilder
	Build() *models.PostStatus
}

// New Builder Initialization
func NewPostStatus() PostStatusBuilder {
	return &postStatusBuilder{}
}

func (pqb *postStatusBuilder) Id(id int) PostStatusBuilder {
	pqb.id = id
	return pqb
}

func (pqb *postStatusBuilder) PostId(postId int) PostStatusBuilder {
	pqb.postId = postId
	return pqb
}

func (pqb *postStatusBuilder) PostTitle(postTitle string) PostStatusBuilder {
	pqb.postTitle = postTitle
	return pqb
}

func (pqb *postStatusBuilder) Status(status int) PostStatusBuilder {
	pqb.status = status
	return pqb
}

func (pqb *postStatusBuilder) Build() *models.PostStatus {
	return &models.PostStatus{
		Id:        pqb.id,
		PostId:    pqb.postId,
		PostTitle: pqb.postTitle,
		Status:    pqb.status,
	}
}
