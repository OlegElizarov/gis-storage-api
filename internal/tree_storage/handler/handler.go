package handler

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gis-storage-api/internal/models"
	"gis-storage-api/internal/tree_storage"
)

type TreeStorage struct {
	usecase tree_storage.TreeUsecase
}

func NewTreeStorage(r *gin.Engine, u tree_storage.TreeUsecase) *TreeStorage {
	storage := &TreeStorage{usecase: u}

	r.GET("/api/v1/get_tree_data", storage.GetTreeDataJson)
	r.POST("/api/v1/add_tree_data", storage.AddTreeDataJson)

	r.GET("/api/v1/get_tree_growth_data", storage.GetTreeGrowthDataJson)
	r.POST("/api/v1/add_tree_growth_data", storage.AddTreeGrowthDataJson)

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
		writeError(c, err)
		return
	}

	err = writeTypedResponse(c, trees)
	if err != nil {
		writeError(c, err)
		return
	}

	return
}

func (ts *TreeStorage) AddTreeDataJson(c *gin.Context) {
	var trees []models.Tree

	err := readTypedRequestData[models.Tree](c, &trees)
	if err != nil {
		writeError(c, err)
		return
	}

	err = ts.usecase.AddTreeData(c, trees)
	if err != nil {
		writeError(c, err)
		return
	}
}

func (ts *TreeStorage) GetTreeGrowthDataJson(c *gin.Context) {
	var selectionRules []string

	if selection := c.Query("selection"); selection != "" && selection != "all" {
		selectionRules = strings.Split(selection, ",")
	}

	var filters models.Filters

	filters.Fill(c.Request.URL.Query())

	trees, err := ts.usecase.GetTreeGrowth(c, models.Selection{Fields: selectionRules}, filters)
	if err != nil {
		writeError(c, err)
		return
	}

	err = writeTypedResponse(c, trees)
	if err != nil {
		writeError(c, err)
		return
	}

	return
}

func (ts *TreeStorage) AddTreeGrowthDataJson(c *gin.Context) {
	trees := []models.GrowthTree{}

	err := readTypedRequestData[models.GrowthTree](c, &trees)
	if err != nil {
		writeError(c, err)
		return
	}

	err = ts.usecase.AddTreeGrowth(c, trees)
	if err != nil {
		writeError(c, err)
		return
	}
}
