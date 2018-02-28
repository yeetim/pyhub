package Cms

import (
	"fmt"

	"github.com/dinever/golf"
	"github.com/eirini/test/app/controllers"
	"github.com/eirini/test/app/models"
)

func Init() {
	dbPath := "root:123456@tcp(192.168.56.110:3306)/gotest?charset=utf8"
	models.Initialize(dbPath)
}

func Run() {
	app := golf.New()
	app = controllers.Initialize(app)
	fmt.Printf("Application Started on port 9000\n")
	app.Run(":9000")
}
