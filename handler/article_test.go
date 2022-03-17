package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"kumparan/handler"
	"kumparan/helper"
	"kumparan/model"
	"kumparan/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/mock"
)

func TestAddArticle(t *testing.T) {
	testCases := []struct {
		Json       map[string]string
		Article    *model.Article
		ErrJson    error
		ErrParam   error
		ErrUsecase error
		StatusCode int
	}{
		{
			Json: map[string]string{"author": "Ibnu", "title": "Judul", "body": "Isi"},
			Article: &model.Article{
				Author: "Ibnu",
				Title:  "Judul",
				Body:   "Isi",
			},
			ErrUsecase: nil,
			StatusCode: http.StatusOK,
		},
		{
			ErrJson:    errors.New("err"),
			StatusCode: http.StatusBadRequest,
		},
		{
			ErrParam:   errors.New("err"),
			StatusCode: http.StatusBadRequest,
		},
		{
			Json:       map[string]string{"author": "Ibnu", "title": "Judul", "body": "Isi"},
			ErrUsecase: errors.New("err"),
			StatusCode: http.StatusInternalServerError,
		},
	}

	gin.SetMode(gin.TestMode)
	helper.GoCache = cache.New(5*time.Minute, 10*time.Minute)

	for _, testCase := range testCases {

		mockUsecase := new(mocks.ArticleUsecase)

		if testCase.ErrParam == nil && testCase.ErrJson == nil {
			mockUsecase.On("Create", mock.AnythingOfType("model.Article")).Return(testCase.Article, testCase.ErrUsecase)
			defer mockUsecase.AssertCalled(t, "Create", mock.AnythingOfType("model.Article"))
		}

		handler := handler.ArticleHandler{
			ArticleUsecase: mockUsecase,
		}

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		jsonValue, _ := json.Marshal(testCase.Json)

		if testCase.ErrJson != nil {
			jsonValue = []byte("")
		}

		c.Request, _ = http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonValue))
		c.Request.Header.Add("Content-Type", "application/json")

		handler.AddArticle(c)

		assert.Equal(t, testCase.StatusCode, rec.Code)
	}
}

func TestViewArticles(t *testing.T) {
	testCases := []struct {
		Articles   *[]model.Article
		ErrUsecase error
		StatusCode int
	}{
		{
			Articles: &[]model.Article{
				model.Article{
					Id:     1,
					Author: "Ibnu",
					Title:  "Judul",
					Body:   "Isi",
				},
			},
			ErrUsecase: nil,
			StatusCode: http.StatusOK,
		},
		{
			ErrUsecase: errors.New("err"),
			StatusCode: http.StatusInternalServerError,
		},
		{
			Articles:   &[]model.Article{},
			StatusCode: http.StatusNotFound,
		},
	}

	gin.SetMode(gin.TestMode)
	helper.GoCache = cache.New(5*time.Minute, 10*time.Minute)

	for _, testCase := range testCases {

		mockUsecase := new(mocks.ArticleUsecase)

		mockUsecase.On("Get", mock.AnythingOfType("model.ViewArticleReq")).Return(testCase.Articles, testCase.ErrUsecase)
		defer mockUsecase.AssertCalled(t, "Get", mock.AnythingOfType("model.ViewArticleReq"))

		handler := handler.ArticleHandler{
			ArticleUsecase: mockUsecase,
		}

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		c.Request, _ = http.NewRequest("GET", "/articles", bytes.NewBuffer([]byte("")))

		handler.ViewArticles(c)

		assert.Equal(t, testCase.StatusCode, rec.Code)
	}
}

func TestNewArticleHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.ArticleUsecase)
	handler.NewArticleHandler(gin.Default(), mockUsecase)
}
