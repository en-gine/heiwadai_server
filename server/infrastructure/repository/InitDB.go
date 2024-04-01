package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"server/infrastructure/env"
	"server/infrastructure/logger"

	_ "github.com/lib/pq"
)

var (
	Conn *sql.DB
	once sync.Once
)

func InitDB() *sql.DB {
	var err error

	if Conn != nil {
		return Conn
	}
	once.Do(func() {
		user := env.GetEnv(env.PsqlUser)
		password := env.GetEnv(env.PsqlPass)
		host := env.GetEnv(env.PsqlHost)
		port := env.GetEnv(env.PsqlPort)
		database := env.GetEnv(env.PsqlDbname)

		Conn, err = sql.Open("postgres",
			fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, database))
		if err != nil {
			logger.Fatalf("OpenError: %v", err)
			panic("DB couldn't be Opened!")
		}

		if err = Conn.Ping(); err != nil {
			logger.Fatalf("PingError: %v", err)
			logger.Warn("DB couldn't be Connected!")
		}
	})

	return Conn
}
