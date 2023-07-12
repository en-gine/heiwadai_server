package implement

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func InitDB() (*sql.DB, error) {
	var err error

	if Conn != nil {
		return Conn, nil
	}

	// err = godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("failed to load .env file: ", err)
	// 	return nil, err
	// }

	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASS")
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	database := os.Getenv("PSQL_DBNAME")

	Conn, err = sql.Open("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	if err != nil {
		log.Fatal("OpenError: ", err)
		return nil, err
	}

	if err = Conn.Ping(); err != nil {
		log.Fatal("PingError: ", err)
		return nil, err
	}

	return Conn, nil
}
