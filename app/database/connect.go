package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     string
	user     string
	password string
	dbname   string
	port     int64
)

func Conn() (*sql.DB, error) {

	err := godotenv.Load("../app/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host = os.Getenv("SQL_HOST")
	user = os.Getenv("SQL_USER")
	password = os.Getenv("SQL_PASSWORD")
	dbname = os.Getenv("SQL_DATABASE")
	port = 5432

	pgsqlConn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pgsqlConn)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
