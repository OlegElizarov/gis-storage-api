package usecase

import (
	"context"
	"gis-storage-api/internal/models"
	"gis-storage-api/internal/tree_storage"
)

type TreeUsecase struct {
	repository tree_storage.TreeRepository
}

func NewTreeUsecase(repository tree_storage.TreeRepository) *TreeUsecase {
	return &TreeUsecase{repository: repository}
}

func (t TreeUsecase) GetTreeData(ctx context.Context, selection models.Selection, filters models.Filters) ([]models.Tree, error) {
	return t.repository.GetTreeData(ctx, selection, filters)
}

func (t TreeUsecase) AddTreeData(ctx context.Context, trees []models.Tree) error {
	return t.repository.AddTreeData(ctx, trees)
}
