package repositories_test

import (
	"testing"

	"github.com/kamadakohei/saki-book/ch3/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kamadakohei/saki-book/ch3/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %d but want %d\n", num, expectedNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ID:        1,
		ArticleID: 1,
		Message:   "test comment1",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.ID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `delete from comments where id = ?`
		testDB.Exec(sqlStr, newComment.ID)
	})
}
