package server

import (
	"gis-storage-api/internal/tree_storage/handler"
	"gis-storage-api/internal/tree_storage/repository"
	"gis-storage-api/internal/tree_storage/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewServer(db *pgx.Conn) *gin.Engine {
	r := gin.Default()
	AddStaticServe(r)

	treeRepository := repository.NewTreeRepository(db)

	treeUsecace := usecase.NewTreeUsecase(treeRepository)

	_ = handler.NewTreeStorage(r, treeUsecace)

	return r
}
