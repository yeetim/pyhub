package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/dinever/golf"
	"github.com/eirini/test/app/models"
)

type signup struct {
	name     string
	email    string
	password string
}

func AuthSignUpGetController(ctx *golf.Context) {
	data := map[string]interface{}{}
	ctx.Loader("admin").Render("index.html", data)
}
func AuthSignUpPostController(ctx *golf.Context) {
	userNum, err := models.GetNumberOfUsers()
	if err != nil || userNum != 0 {
		ctx.Abort(403)
		return
	}
	name := ctx.Request.FormValue("name")
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(password)

}

func AuthLoginGetController(ctx *golf.Context) {
	data := map[string]interface{}{}
	ctx.Loader("admin").Render("index.html", data)
}

func AuthLoginPostController(ctx *golf.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	fmt.Println(decoder)
	var t signup
	err := decoder.Decode(&t)
	if err != nil {
		ctx.Abort(403)
	}
	fmt.Println(t)
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")
	user := &models.User{Email: email}
	//err := user.GetUserByEmail()
	if user == nil || err != nil {
		ctx.Send("error")
	}
	fmt.Println(email)
	fmt.Println(password)
}
