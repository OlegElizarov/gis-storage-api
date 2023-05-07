package models

import (
	"github.com/doug-martin/goqu/v9"
	"net/url"
	"strings"
)

type Filter struct {
	Name      string
	Separator string
	keyType   string
	Values    []string
}

func (f Filter) Apply(sql *goqu.SelectDataset) *goqu.SelectDataset {
	switch f.keyType {
	case boolType:
		if len(f.Values) > 1 {
			return nil
		}
		return sql.Where(
			goqu.C(f.Name).Eq(f.Values[0]),
		)
	case timestampType:
		if len(f.Values) == 2 {
			if f.Values[0] == f.Values[1] {
				return sql.Where(
					goqu.C(f.Name).Eq(f.Values[0]),
				)
			}
			return sql.
				Where(
					goqu.C(f.Name).Gte(f.Values[0]),
					goqu.C(f.Name).Lte(f.Values[1]),
				)
		}

		values := make([]interface{}, len(f.Values))
		for ind := range f.Values {
			values[ind] = f.Values[ind]
		}

		return sql.Where(
			goqu.C(f.Name).In(values...),
		)

	case stringType:
		values := make([]interface{}, len(f.Values))
		for ind := range f.Values {
			values[ind] = f.Values[ind]
		}

		return sql.Where(
			goqu.Or(
				goqu.C(f.Name).In(values...),
			),
		)

	case intType, floatType:
		if len(f.Values) == 2 {
			if f.Values[0] == f.Values[1] {
				return sql.Where(
					goqu.C(f.Name).Eq(f.Values[0]),
				)
			}
			return sql.
				Where(
					goqu.C(f.Name).Gte(f.Values[0]),
					goqu.C(f.Name).Lte(f.Values[1]),
				)
		}

		values := make([]interface{}, len(f.Values))
		for ind := range f.Values {
			values[ind] = f.Values[ind]
		}

		return sql.Where(
			goqu.C(f.Name).In(values...),
		)

	}

	return nil
}

type Filters struct {
	Data []Filter
}

func (ff *Filters) Fill(query url.Values) {
	if ff == nil {
		ff.Data = make([]Filter, 0, len(query))
	}
	for key := range query {
		if field, ok := fieldsMap[key]; ok {
			ff.Data = append(ff.Data, Filter{
				Name:      key,
				Separator: field.Separator,
				Values:    strings.Split(query[key][0], field.Separator),
				keyType:   field.Type,
			})
		}
	}
}

var fieldsMap = map[string]struct {
	Separator string
	Type      string
}{
	// tree_data
	//"id":       {Separator: ";", Type: intType},
	"gis_id":           {Separator: ";", Type: intType},
	"pcd_name":         {Separator: ",", Type: stringType},
	"x_coordinate":     {Separator: ";", Type: floatType},
	"y_coordinate":     {Separator: ";", Type: floatType},
	"gis_height_mitro": {Separator: ";", Type: floatType},
	"gis_height_il":    {Separator: ";", Type: floatType},
	"order_number":     {Separator: ";", Type: intType},
	"tree_type":        {Separator: ",", Type: stringType},
	"circle":           {Separator: ";", Type: floatType},
	"diameter_mitro":   {Separator: ";", Type: floatType},
	"diameter_il":      {Separator: ";", Type: floatType},
	// tree_growth

	//"gis_id":   {Separator: ";", Type: intType},
	//"id":       {Separator: ";", Type: intType},
	"ts":       {Separator: ";", Type: timestampType},
	"age":      {Separator: ";", Type: intType},
	"diameter": {Separator: ";", Type: floatType},
	"height":   {Separator: ";", Type: floatType},
	"is_alive": {Separator: ";", Type: boolType},
}

const (
	intType       = "int64"
	floatType     = "float64"
	stringType    = "string"
	boolType      = "bool"
	timestampType = "ts"
)
