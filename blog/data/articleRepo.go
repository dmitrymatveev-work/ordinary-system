package data

import (
	"errors"
	"ordinary-system/blog/model"
)

func CreateArticle(userID int64, article model.Article) (model.Article, error) {
	return model.Article{}, errors.New("not implemented")
}

func GetArticles(userID int64) ([]model.Article, error) {
	return nil, errors.New("not implemented")
}
