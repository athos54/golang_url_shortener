package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var Db *sql.DB
var err error

func Connect() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", user, pass, dbname)
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func SearchURL(us *UrlShorten) {
	rows, err := Db.Query("SELECT long_url FROM urls where short_path = $1", us.ShortPath)
	if err != nil {
		log.Error(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&us.LongURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(us)
	}
}

func AppendURL(us *UrlShorten) error {
	index := 0

	sqlStatement := `
		INSERT INTO urls (
			"short_path", 
			"long_url"
		)
		VALUES ($1, $2)
		RETURNING id
	`
	err := Db.QueryRow(
		sqlStatement,
		us.ShortPath,
		us.LongURL,
	).Scan(&index)

	if err != nil {
		return err
	}

	return nil
}
