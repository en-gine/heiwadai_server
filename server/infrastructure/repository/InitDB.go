package repository

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"server/infrastructure/env"
	"server/infrastructure/logger"

	_ "github.com/lib/pq"
)

var (
	conn *sql.DB  // 小文字にして外部から直接アクセスを防ぐ
	once sync.Once
	mu   sync.RWMutex  // 読み取り時の競合状態を防ぐ
)

func InitDB() *sql.DB {
	mu.RLock()
	if conn != nil {
		mu.RUnlock()
		return conn
	}
	mu.RUnlock()

	var err error
	once.Do(func() {
		user := env.GetEnv(env.PsqlUser)
		password := env.GetEnv(env.PsqlPass)
		host := env.GetEnv(env.PsqlHost)
		port := env.GetEnv(env.PsqlPort)
		database := env.GetEnv(env.PsqlDbname)

		mu.Lock()
		conn, err = sql.Open("postgres",
			fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo", user, password, host, port, database))
		mu.Unlock()
		
		if err != nil {
			logger.Fatalf("OpenError: %v", err)
			panic("DB couldn't be Opened!")
		}

		if err = conn.Ping(); err != nil {
			logger.Fatalf("PingError: %v", err)
			logger.Warn("DB couldn't be Connected!")
		}

		// Configure connection pool to prevent connection exhaustion
		// Supabase typically allows 60 connections per project in free tier
		conn.SetMaxOpenConns(20)              // Maximum number of open connections
		conn.SetMaxIdleConns(5)               // Maximum number of idle connections
		conn.SetConnMaxLifetime(5 * time.Minute)  // Maximum lifetime of a connection
		conn.SetConnMaxIdleTime(1 * time.Minute)  // Maximum idle time for connections
		
		logger.Info("Database connection pool configured: MaxOpenConns=20, MaxIdleConns=5")
	})

	return conn
}

// func SetDebugSQL(ctx *context.Context) *context.Context {
// 	boil.SetDB(Conn)
// 	sqlBoilerDebugCtx := boil.WithDebug(*ctx, env.GetEnv(env.EnvMode) == "dev")
// 	sqlBoilerLoggerDebugCtx := boil.WithDebugWriter(sqlBoilerDebugCtx, os.Stderr)
// 	return &sqlBoilerLoggerDebugCtx
// }
