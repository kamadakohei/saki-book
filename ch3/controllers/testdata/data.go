package testdata

import "github.com/kamadakohei/saki-book/ch3/models"

var articleTestData = []models.Article{
	{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "saki",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  4,
	},
}

var commentTestData = []models.Comment{
	{
		ID:        1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	{
		ID:        2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
