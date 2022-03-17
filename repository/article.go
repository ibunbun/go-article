package repository

import (
	"fmt"

	"kumparan/model"

	"github.com/jinzhu/gorm"
)

type Article struct {
	DB *gorm.DB
}

type ArticleRepo interface {
	Create(article model.Article) (*model.Article, error)
	Get(params model.ViewArticleReq) (*[]model.Article, error)
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &Article{DB: db}
}

func (a *Article) Create(article model.Article) (*model.Article, error) {
	err := a.DB.Save(&article).Error
	if err != nil {
		fmt.Printf("[ArticleRepo.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("Server error")
	}
	return &article, nil
}

func (a *Article) Get(params model.ViewArticleReq) (*[]model.Article, error) {
	var articles []model.Article

	query := a.DB.Table("articles")

	if params.Author != "" {
		author := "%" + params.Author + "%"
		query = query.Where("author LIKE ?", author)
	}

	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		query = query.Where("title LIKE ? OR body LIKE ?", keyword, keyword)
	}

	err := query.Order("created desc").Find(&articles).Error
	if err != nil {
		fmt.Printf("[ArticleRepo.Get] error execute query %v \n", err)
		return nil, fmt.Errorf("Server error")
	}

	return &articles, nil
}
