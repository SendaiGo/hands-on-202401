package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// データベースに接続
	db, err := connectSQLite3()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close() // main関数終了時にデータベース接続をクローズ

	var id int
	var name string

	// データベースクエリ
	u := db.QueryRow("SELECT id,name FROM users WHERE id = ?", 1)

	// ユーザー情報の取得
	if err := u.Scan(&id, &name); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %d, name: %s\n", id, name) // 取得したユーザー情報の出力
}

// SQlite3の接続
func connectSQLite3() (*sql.DB, error) {
	// データベースのファイルパスの設定
	dbPath := "data/sqlite3.db" // 適宜パスを指定してください

	// データベースのオープン
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// db.Close()の呼び出しはmain関数内でdeferを使って行う
	return db, nil
}
