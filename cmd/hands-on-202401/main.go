package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// データベースのファイルパスの設定
	dbPath := "data/sqlite3.db"

	// データベースのオープン
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.File("public/index.html")
	})

	// TODOのCRUD処理
	// Create
	e.POST("/todo", func(c echo.Context) error {
		return create(c, db)
	})

	// ReadAll
	e.GET("/todo", func(c echo.Context) error {
		return readAll(c, db)
	})
	// Read
	e.GET("/todo/:id", func(c echo.Context) error {
		return read(c, db)
	})

	// Update
	e.PUT("/todo/:id", func(c echo.Context) error {
		return update(c, db)
	})

	// Delete
	e.DELETE("/todo/:id", func(c echo.Context) error {
		return delete(c, db)
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":80"))
}
