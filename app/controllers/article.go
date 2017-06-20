package controllers

import (
	"PT-Blog/app/models"
	"github.com/revel/revel"
)

type Article struct {
	*revel.Controller
}

func (c Article) GetArticleList() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(err)
	}
	defer dao.Close()
	articleList := dao.FindArticles()
	return c.RenderJSON(articleList)
}

func (c Article) GetArticle() revel.Result {
	id := c.Params.Route.Get("id")
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(err)
	}
	defer dao.Close()
	article := dao.FindArticleById(id)
	return c.RenderJSON(article)
}
