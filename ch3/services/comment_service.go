package services

import (
	"github.com/kamadakohei/saki-book/ch3/apperrors"
	"github.com/kamadakohei/saki-book/ch3/models"
	"github.com/kamadakohei/saki-book/ch3/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to insert comment")
		return models.Comment{}, err
	}

	return newComment, nil
}
