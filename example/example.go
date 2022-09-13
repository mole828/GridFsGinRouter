package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	GridGin "github.com/mole828/GridFsGinRouter"
	"gopkg.in/mgo.v2"
)

func main() {
	dial, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return
	}
	db := dial.DB("moles")
	app := gin.New()
	app.LoadHTMLGlob("template/**")
	app.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{})
	})
	app.Use(gin.Logger())

	group := app.Group("/")

	GridGin.ServeGroup(group, db)

	port := 7999
	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("listen: %v", addr)
	if err := app.Run(fmt.Sprintf(":%d", port)); err != nil {
		return
	}
}
