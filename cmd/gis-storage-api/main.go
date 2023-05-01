package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"gis-storage-api/internal/server"
	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"io"
)

func main() {
	ctx := context.Background()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})

	db, err := InitDatabase(ctx)
	if err != nil {
		fmt.Println("QQQ", err)
		return
	}

	r := server.NewServer(db)
	err = r.Run()
	if err != nil {
		return
	}
}

const dsn = `host=localhost port=5432 user=gis password=gis dbname=gis_storage sslmode=disable`

func InitDatabase(ctx context.Context) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return pgx.ConnectConfig(ctx, config)
}
