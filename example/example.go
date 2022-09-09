package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyahm/golog"
	"gopkg.in/mgo.v2"
)

func main() {
	dial, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return
	}
	db := dial.DB("moles")
	app := gin.New()
	app.LoadHTMLGlob("template/*")
	app.Use(gin.Logger())
	ServeGroup(&app.RouterGroup, db)

	port := 7999
	addr := fmt.Sprintf("localhost:%d", port)
	golog.Infof("listen: %v", addr)
	if err := app.Run(fmt.Sprintf(":%d", port)); err != nil {
		return
	}
}
