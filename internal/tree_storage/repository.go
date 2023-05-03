package tree_storage

import (
	"context"
	"gis-storage-api/internal/models"
)

type TreeRepository interface {
	GetTreeData(context.Context, models.Selection, models.Filters) ([]models.Tree, error)
	AddTreeData(context.Context, []models.Tree) error

	GetTreeGrowth(context.Context, models.Selection, models.Filters) ([]models.Tree, error)
	AddTreeGrowth(context.Context, []models.Tree) error
}
