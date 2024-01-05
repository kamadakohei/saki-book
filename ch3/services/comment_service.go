package services

import (
	"github.com/kamadakohei/saki-book/ch3/models"
	"github.com/kamadakohei/saki-book/ch3/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
