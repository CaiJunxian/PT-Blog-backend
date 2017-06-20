package controllers

import (
	"PT-Blog/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Test() revel.Result {
	result := models.Result{"0", "success"}
	return c.RenderJSON(result)
}
