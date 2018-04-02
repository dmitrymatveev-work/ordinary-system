package data

import (
	"fmt"
	"ordinary-system/blog/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateArticle function stores article into a storage
func CreateArticle(userID int64, article model.Article) (model.Article, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return model.Article{}, err
	}
	defer session.Close()

	c := session.DB(fmt.Sprintf("user%dBlog", userID)).C("articles")
	err = c.Insert(&article)
	if err != nil {
		return model.Article{}, err
	}

	return article, nil
}

// GetArticles function retrieves articles from a storage
func GetArticles(userID int64) ([]model.Article, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(fmt.Sprintf("user%dBlog", userID)).C("articles")
	var articles []model.Article
	err = c.Find(bson.M{}).All(&articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
