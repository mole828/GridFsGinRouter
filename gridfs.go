package GridGin

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ServeGroup(app *gin.RouterGroup, db *mgo.Database) {

	bucket := db.GridFS("fs")

	app.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{})
	})

	app.GET("/list", func(context *gin.Context) {
		find := bucket.Find(nil)
		var result []interface{}
		find.All(&result)
		context.JSON(200, result)
	})

	app.POST("/set", func(context *gin.Context) {
		file, m, err := context.Request.FormFile("file")
		if err != nil {
			return
		}

		create, err := bucket.Create(m.Filename)
		if err != nil {
			return
		}
		defer func() {
			err := create.Close()
			if err != nil {
				log.Printf(err.Error())
				return
			}
		}()

		bytes, err := io.ReadAll(file)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err = create.Write(bytes)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"_id": create.Id(),
		})
	})

	app.GET("/get/:key", func(context *gin.Context) {
		key := context.Param("key")
		if !bson.IsObjectIdHex(key) {
			context.JSON(404, gin.H{"msg": "can't accept key"})
			return
		}
		_id := bson.ObjectIdHex(key)
		file, err := bucket.OpenId(_id)
		if err != nil {
			context.JSON(404, gin.H{"msg": err.Error()})
			return
		}
		http.ServeContent(context.Writer, context.Request, file.Name(), file.UploadDate(), file)
	})
}
