package builder

import "github.com/wincentrtz/fake-news/models"

type postQueueBuilder struct {
	id        int
	postId    int
	postTitle string
	progress  int
}

// PostQueueBuilder builder interface
type PostQueueBuilder interface {
	Id(int) PostQueueBuilder
	PostId(int) PostQueueBuilder
	PostTitle(string) PostQueueBuilder
	Progress(int) PostQueueBuilder
	Build() *models.PostQueue
}

// New Builder Initialization
func NewPostQueue() PostQueueBuilder {
	return &postQueueBuilder{}
}

func (pqb *postQueueBuilder) Id(id int) PostQueueBuilder {
	pqb.id = id
	return pqb
}

func (pqb *postQueueBuilder) PostId(postId int) PostQueueBuilder {
	pqb.postId = postId
	return pqb
}

func (pqb *postQueueBuilder) PostTitle(postTitle string) PostQueueBuilder {
	pqb.postTitle = postTitle
	return pqb
}

func (pqb *postQueueBuilder) Progress(progress int) PostQueueBuilder {
	pqb.progress = progress
	return pqb
}

func (pqb *postQueueBuilder) Build() *models.PostQueue {
	return &models.PostQueue{
		Id:        pqb.id,
		PostId:    pqb.postId,
		PostTitle: pqb.postTitle,
		Progress:  pqb.progress,
	}
}
