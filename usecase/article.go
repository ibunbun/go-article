package usecase

import (
	"kumparan/model"
	"kumparan/repository"
)

type Article struct {
	ArticleRepo repository.ArticleRepo
}

type ArticleUsecase interface {
	Create(article model.Article) (*model.Article, error)
	Get(params model.ViewArticleReq) (*[]model.Article, error)
}

func NewArticleUsecase(articleRepo repository.ArticleRepo) ArticleUsecase {
	return &Article{articleRepo}
}

func (a *Article) Create(article model.Article) (*model.Article, error) {
	return a.ArticleRepo.Create(article)
}

func (a *Article) Get(params model.ViewArticleReq) (*[]model.Article, error) {
	return a.ArticleRepo.Get(params)
}
