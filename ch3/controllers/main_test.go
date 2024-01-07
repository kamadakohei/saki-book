package controllers_test

import (
	"testing"

	"github.com/kamadakohei/saki-book/ch3/controllers/testdata"

	"github.com/kamadakohei/saki-book/ch3/controllers"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
