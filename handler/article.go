package handler

import (
	"net/http"
	"time"

	"kumparan/helper"
	"kumparan/model"
	"kumparan/usecase"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	ArticleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(r *gin.Engine, articleUsecase usecase.ArticleUsecase) {
	articleHandler := ArticleHandler{articleUsecase}

	r.POST("/articles", articleHandler.AddArticle)
	r.GET("/articles", articleHandler.ViewArticles)
}

func (ah *ArticleHandler) AddArticle(c *gin.Context) {
	var req model.AddArticleReq
	err := c.BindJSON(&req)
	if err != nil {
		helper.HandleError(c, http.StatusBadRequest, "Something is wrong")
		return
	}

	if req.Author == "" || req.Title == "" || req.Body == "" {
		helper.HandleError(c, http.StatusBadRequest, "Missing required parameter")
		return
	}

	article, err := ah.ArticleUsecase.Create(model.Article{
		Author:  req.Author,
		Title:   req.Title,
		Body:    req.Body,
		Created: time.Now(),
	})

	if err != nil {
		helper.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.GoCache.Flush()

	helper.HandleSuccess(c, article)
	return
}

func (ah *ArticleHandler) ViewArticles(c *gin.Context) {
	req := model.ViewArticleReq{
		Author:  helper.RemoveNonAlphaNumSpace(c.Query("author")),
		Keyword: helper.RemoveNonAlphaNumSpace(c.Query("query")),
	}

	articles, err := ah.ArticleUsecase.Get(req)
	if err != nil {
		helper.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*articles) == 0 {
		helper.HandleError(c, http.StatusNotFound, "Article not found")
		return
	}

	url := helper.GetFullReqUrl(c)
	helper.GoCache.Set(url, articles, 0)

	helper.HandleSuccess(c, articles)
	return
}
