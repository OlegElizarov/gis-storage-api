package repository

import (
	"context"
	"fmt"
	"gis-storage-api/internal/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"strings"
)

var allFields = "gis_id, pcd_name, x_coordinate, y_coordinate, gis_height_mitro, gis_height_il, order_number, tree_type, circle, diameter_mitro, diameter_il"
var allFieldsColumns = strings.Split(allFields, ", ")

type TreeRepository struct {
	db *pgx.Conn
}

func NewTreeRepository(db *pgx.Conn) *TreeRepository {
	return &TreeRepository{db: db}
}

func (tr *TreeRepository) GetTreeData(ctx context.Context, selection models.Selection, filters models.Filters) ([]models.Tree, error) {
	var fields []interface{}
	var source []string
	if len(selection.Fields) != 0 {
		source = selection.Fields
	} else {
		source = allFieldsColumns
	}

	for ind := range source {
		fields = append(fields, source[ind])
	}

	sql := goqu.From("tree_data").
		Select(
			fields...,
		)

	for ind := range filters.Data {
		filter := filters.Data[ind]

		sql = filter.Apply(sql)
	}

	sqlRaw, _, _ := sql.ToSQL()

	row, err := tr.db.Query(context.Background(), sqlRaw)
	if err != nil {
		return nil, err
	}

	trees := []models.Tree{}
	err = pgxscan.ScanAll(&trees, row)
	if err != nil {
		return nil, err
	}
	return trees, nil
}

func (tr *TreeRepository) AddTreeData(ctx context.Context, trees []models.Tree) error {
	sqlStr := `INSERT INTO tree_data(` + allFields + `) VALUES`

	for _, tree := range trees {
		sqlStr += fmt.Sprintf("(%d,'%s',%v,%v,%v,%v,%d,'%s',%v,%v,%v),",
			tree.GisID, tree.PcdName, tree.XCoordinate, tree.YCoordinate,
			tree.GISHeightMitro, tree.GISHeightIl,
			tree.OrderNumber, tree.TreeType, tree.Circle,
			tree.DiameterMitro, tree.DiameterIl,
		)
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	sqlStr += " on conflict do nothing"

	//format all vals at once
	_, err := tr.db.Exec(ctx, sqlStr)
	return err
}