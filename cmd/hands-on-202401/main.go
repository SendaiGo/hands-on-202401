package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// データベースのファイルパスの設定
	dbPath := "../../data/sqlite3.db"

	// データベースのオープン
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// ハンドラの登録
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slice := make([]string, 0, 10)
		for _, v := range []string{"a", "b", "c"} {
			slice = append(slice, fmt.Sprintf("%d", v))
		}
		fmt.Fprint(w, "Hello, World!")
	})

	// サーバーの起動
	log.Println("start http server :80")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
