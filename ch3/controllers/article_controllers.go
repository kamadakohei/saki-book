package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kamadakohei/saki-book/ch3/apperrors"

	"github.com/gorilla/mux"
	"github.com/kamadakohei/saki-book/ch3/controllers/services"
	"github.com/kamadakohei/saki-book/ch3/models"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{
		service: s,
	}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "query parameter must be integer")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "path param must be integer")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}
