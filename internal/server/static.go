package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStaticServe(r *gin.Engine) {
	r.StaticFS("/static", http.Dir("static"))

	r.POST("/api/v1/upload", func(c *gin.Context) {
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			writeError(c, err)
		}
		files := form.File["data"]

		dirToSave := c.Query("dir")
		finalPath := "static/" + dirToSave + "/"

		for _, file := range files {
			fileName := finalPath + file.Filename
			log.Println(fileName)

			// TODO: check or sanitize input here

			err := c.SaveUploadedFile(file, fileName)
			if err != nil {
				writeError(c, err)
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
}

func writeError(c *gin.Context, err error) {
	c.Writer.WriteHeader(500)
	c.Writer.Write([]byte("error: " + err.Error()))
}
