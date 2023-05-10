package handler

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

const (
	jsonType = "json"
	csvType  = "csv"
)

func writeError(c *gin.Context, err error) {
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

func readTypedRequestData[T comparable](c *gin.Context, dst *[]T) error {
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
	unmarshalFunc := getUnmarshalFunc(file)
	if unmarshalFunc == nil {
		return errors.New("wrong input file(extension or content-type)")
	}

	err = unmarshalFunc(fileData, dst)
	if err != nil {
		return err
	}

	return nil
}

func getUnmarshalFunc(file *multipart.FileHeader) func([]byte, interface{}) error {
	dataType := file.Header.Get("Content-Type")

	switch dataType {
	case "text/csv":
		return gocsv.UnmarshalBytes
	case "application/json":
		return json.Unmarshal
	}

	extension := path.Ext(file.Filename)
	switch extension {
	case ".csv":
		return gocsv.UnmarshalBytes
	case ".json":
		return json.Unmarshal
	}

	return nil
}
