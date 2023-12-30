package repositories

import (
	"database/sql"

	"github.com/kamadakohei/saki-book/ch3/models"
)

func InsertComment(db *sql.DB, article models.Comment) (models.Comment, error) {
	const sqlStr = `INSERT INTO comments (article_id, message, created_at) VALUES (?, ?, now())`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = article.ArticleID, article.Message

	result, err := db.Exec(sqlStr, article.ArticleID, article.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.ID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `SELECT * FROM comments where article_id = ?;`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.ID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
