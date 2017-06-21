package controllers

import (
	"PT-Blog/app/models"
	"github.com/revel/revel"
	"strings"
)

type Article struct {
	App
}

// get article list
func (c Article) GetArticleList() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		revel.ERROR.Printf("%s", err)
		c.Response.Status = 500
		return c.RenderJSON(models.Error("article.err.1", "internal error"))
	}
	defer dao.Close()
	articleList := dao.FindArticles()
	return c.RenderJSON(models.Success(articleList))
}

// get a article by id
func (c Article) GetArticle() revel.Result {
	id := c.Params.Route.Get("id")
	dao, err := models.NewDao()
	if err != nil {
		revel.ERROR.Printf("%s", err)
		c.Response.Status = 500
		return c.RenderJSON(models.Error("article.err.1", "internal error"))
	}
	defer dao.Close()
	article := dao.FindArticleById(id)
	return c.RenderJSON(models.Success(article))
}

func (c Article) CreateArticle(article *models.Article) revel.Result {
	article.Title = strings.TrimSpace(article.Title)
	article.Email = strings.TrimSpace(article.Email)
	article.Subject = strings.TrimSpace(article.Subject)
	article.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.RenderJSON(models.Error("article.err.2", "validation error"))
	}
	dao, err := models.NewDao()
	if err != nil {
		revel.ERROR.Printf("%s", err)
		c.Response.Status = 500
		return c.RenderJSON(models.Error("article.err.1", "internal error"))
	}
	defer dao.Close()
	err = dao.CreateArticle(article)
	if err != nil {
		revel.ERROR.Printf("%s", err)
		c.Response.Status = 500
		return c.RenderJSON(models.Error("article.err.3", "internal error"))
	}
	return c.RenderJSON(models.Success(article))
}
