package main

import (
	"log"
	"time"

	"kumparan/config"
	"kumparan/handler"
	"kumparan/helper"
	"kumparan/repository"
	"kumparan/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickmn/go-cache"
)

func main() {
	db := config.DbConnect()
	defer db.Close()

	helper.GoCache = cache.New(5*time.Minute, 10*time.Minute)

	router := gin.Default()
	router.Use(helper.CheckCache)

	articleRepo := repository.NewArticleRepo(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepo)

	handler.NewArticleHandler(router, articleUsecase)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
