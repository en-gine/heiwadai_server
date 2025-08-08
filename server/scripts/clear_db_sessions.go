package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// 環境変数から接続情報を取得
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASS"),
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_DBNAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	// 現在のセッションを表示
	fmt.Println("Current sessions:")
	rows, err := db.Query(`
		SELECT pid, application_name, client_addr, state, query_start, state_change
		FROM pg_stat_activity 
		WHERE datname = current_database() 
		AND pid <> pg_backend_pid()
		ORDER BY query_start DESC
	`)
	if err != nil {
		log.Fatal("Failed to query sessions:", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var pid int
		var appName, clientAddr, state sql.NullString
		var queryStart, stateChange sql.NullTime
		
		err := rows.Scan(&pid, &appName, &clientAddr, &state, &queryStart, &stateChange)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		
		fmt.Printf("PID: %d, App: %s, Client: %s, State: %s\n", 
			pid, appName.String, clientAddr.String, state.String)
		count++
	}

	if count == 0 {
		fmt.Println("No active sessions found.")
		return
	}

	// ユーザーに確認
	fmt.Printf("\nFound %d active sessions. Clear all? (yes/no): ", count)
	var response string
	fmt.Scanln(&response)

	if response != "yes" {
		fmt.Println("Cancelled.")
		return
	}

	// セッションをクリア
	result, err := db.Exec(`
		SELECT pg_terminate_backend(pid) 
		FROM pg_stat_activity 
		WHERE datname = current_database() 
		AND pid <> pg_backend_pid()
		AND application_name != 'supabase_admin'
	`)
	if err != nil {
		log.Fatal("Failed to terminate sessions:", err)
	}

	affected, _ := result.RowsAffected()
	fmt.Printf("Terminated %d sessions.\n", affected)
}