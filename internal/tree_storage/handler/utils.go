package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

const (
	jsonType = "json"
	csvType  = "csv"
)

func writeError(c *gin.Context, err error) {
	fmt.Println(err)
	c.Writer.WriteHeader(500)
	c.Writer.Write([]byte("error: " + err.Error()))
}

func writeTypedResponse(c *gin.Context, data any) error {
	resultType := c.Query("type")
	if resultType == "" {
		resultType = jsonType
	}

	switch resultType {
	case csvType:
		data, err := gocsv.MarshalBytes(data)
		if err != nil {
			return err
		}
		_, err = c.Writer.Write(data)
		if err != nil {
			return err
		}
	case jsonType:
		c.JSON(http.StatusOK, data)
	}

	return nil
}

func readTypedRequestData(c *gin.Context, dst any) error {
	file, err := c.FormFile("data")
	if err != nil {
		return err
	}
	ff, err := file.Open()
	defer ff.Close()
	if err != nil {
		return nil
	}
	fileData, err := io.ReadAll(ff)
	if err != nil {
		return nil
	}
	dataType := file.Header.Get("Content-Type")

	switch dataType {
	case "text/csv":
		err = gocsv.UnmarshalBytes(fileData, &dst)
	case "application/json":
		err = json.Unmarshal(fileData, &dst)
	}
	if err != nil {
		return err
	}

	return nil
}