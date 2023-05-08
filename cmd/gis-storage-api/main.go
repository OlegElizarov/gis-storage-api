package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"

	"gis-storage-api/internal/server"
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
		log.Println("failed to init to db:", err)
		return
	}

	r := server.NewServer(db)
	err = r.Run()
	if err != nil {
		log.Println("failed to init to app:", err)
		return
	}
}

const dsn = `host=localhost port=5432 user=gis password=gis dbname=gis_storage sslmode=disable`

func InitDatabase(ctx context.Context) (*pgx.Conn, error) {
	log.Println("Initing database")
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	return pgx.ConnectConfig(ctx, config)
}
