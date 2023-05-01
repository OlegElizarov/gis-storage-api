package handler

import (
	"encoding/json"
	"fmt"
	"gis-storage-api/internal/models"
	"github.com/gocarina/gocsv"
	"io"
	"net/http"
	"strings"

	"gis-storage-api/internal/tree_storage"
	"github.com/gin-gonic/gin"
)

const (
	jsonType = "json"
	csvType  = "csv"
)

type TreeStorage struct {
	usecase tree_storage.TreeUsecase
}

func NewTreeStorage(r *gin.Engine, u tree_storage.TreeUsecase) *TreeStorage {
	storage := &TreeStorage{usecase: u}

	r.GET("/api/v1/get_tree_data", storage.GetTreeDataJson)
	r.POST("/api/v1/add_tree_data", storage.AddTreeDataJson)

	return storage
}

func (ts *TreeStorage) GetTreeDataJson(c *gin.Context) {
	var selectionRules []string

	if selection := c.Query("selection"); selection != "" && selection != "all" {
		selectionRules = strings.Split(selection, ",")
	}

	var filters models.Filters

	filters.Fill(c.Request.URL.Query())

	trees, err := ts.usecase.GetTreeData(c, models.Selection{Fields: selectionRules}, filters)
	if err != nil {
		WriteError(c, err)
		return
	}

	resultType := c.Query("type")
	if resultType == "" {
		resultType = jsonType
	}

	switch resultType {
	case csvType:
		data, err := gocsv.MarshalBytes(trees)
		if err != nil {
			WriteError(c, err)
			return
		}
		_, err = c.Writer.Write(data)
		if err != nil {
			WriteError(c, err)
			return
		}
	case jsonType:
		c.JSON(http.StatusOK, trees)
	}
	return
}

func (ts *TreeStorage) AddTreeDataJson(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		WriteError(c, err)
		return
	}
	ff, err := file.Open()
	defer ff.Close()
	if err != nil {
		WriteError(c, err)
		return
	}
	fileData, err := io.ReadAll(ff)
	if err != nil {
		WriteError(c, err)
		return
	}
	dataType := file.Header.Get("Content-Type")
	var trees []models.Tree

	switch dataType {
	case "text/csv":
		err = gocsv.UnmarshalBytes(fileData, &trees)
	case "application/json":
		err = json.Unmarshal(fileData, &trees)
	}
	if err != nil {
		WriteError(c, err)
		return
	}

	err = ts.usecase.AddTreeData(c, trees)
	if err != nil {
		WriteError(c, err)
		return
	}
}

func WriteError(c *gin.Context, err error) {
	fmt.Println(err)
	c.Writer.WriteHeader(500)
	c.Writer.Write([]byte("error: " + err.Error()))
}
