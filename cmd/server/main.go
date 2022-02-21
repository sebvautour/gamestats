package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sebvautour/gamestats/internal/api"
	"github.com/sebvautour/gamestats/internal/webserver"
)

const defaultSqliteDB = "./gamestats.db"

func main() {
	ctx := context.Background()

	sqliteDB := os.Getenv("GAMESTATS_SQLITE_DB")
	if sqliteDB == "" {
		sqliteDB = defaultSqliteDB
	}

	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gameStatsAPI, err := api.New(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(webserver.New(gameStatsAPI).Run(":8080"))
}
