package repository

import (
	"testing"
	"time"
)

func TestDatabaseTimezone(t *testing.T) {
	// データベース接続を初期化
	db := InitDB()
	if db == nil {
		t.Fatal("Failed to initialize database")
	}

	// データベースの現在時刻とタイムゾーンを確認
	var dbTime time.Time
	var timezone string
	err := db.QueryRow("SELECT NOW(), current_setting('TIMEZONE')").Scan(&dbTime, &timezone)
	if err != nil {
		t.Fatalf("Failed to query database time: %v", err)
	}

	// サーバーの現在時刻を取得
	serverTime := time.Now()

	// タイムゾーンが Asia/Tokyo に設定されているか確認
	if timezone != "Asia/Tokyo" {
		t.Errorf("Expected timezone 'Asia/Tokyo', got '%s'", timezone)
	}

	// データベース時刻とサーバー時刻の差が1秒以内であることを確認
	diff := serverTime.Sub(dbTime).Abs()
	if diff > time.Second {
		t.Errorf("Time difference between server and database is too large: %v", diff)
		t.Logf("Server time: %v", serverTime)
		t.Logf("Database time: %v", dbTime)
	}

	// JSTタイムゾーンでの時刻フォーマットを確認
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t.Logf("Database timezone: %s", timezone)
	t.Logf("Database time (JST): %s", dbTime.In(jst).Format("2006-01-02 15:04:05 MST"))
	t.Logf("Server time (JST): %s", serverTime.In(jst).Format("2006-01-02 15:04:05 MST"))
}