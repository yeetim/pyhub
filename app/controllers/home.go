package controllers

import (
	"github.com/dinever/golf"
)

func HomeController(ctx *golf.Context) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	ctx.Loader("theme").Render("index.html", data)
}
