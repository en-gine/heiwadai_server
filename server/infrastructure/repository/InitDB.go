package repository

import (
	"database/sql"
	"fmt"
	"os"

	"server/infrastructure/logger"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func InitDB() *sql.DB {
	var err error

	if Conn != nil {
		return Conn
	}

	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASS")
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	database := os.Getenv("PSQL_DBNAME")

	Conn, err = sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, database))
	if err != nil {
		logger.Fatalf("OpenError: %v", err)
		panic("DB couldn't be Opened!")
	}

	if err = Conn.Ping(); err != nil {
		logger.Fatalf("PingError: %v", err)
		panic("DB couldn't be Connected!")
	}

	return Conn
}
