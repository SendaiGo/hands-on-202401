package main

import (
	"crypto/subtle"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// // ハンドラの登録
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	slice := make([]string, 0, 10)
	// 	for _, v := range []string{"a", "b", "c"} {
	// 		slice = append(slice, fmt.Sprintf("%d", v))
	// 	}
	// 	fmt.Fprint(w, "Hello, World!")
	// })

	// // サーバーの起動
	// log.Println("start http server :80")
	// err = http.ListenAndServe(":80", nil)
	// if err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// Echo による実装
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		// レスポンスの返却
		// return c.String(http.StatusOK, "Hello, World!")

		// JSON を返す場合
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	e.POST("/hello", func(c echo.Context) error {
		// リクエストボディの取得

		// curl -X POST -H "Content-Type: application/json" -d '{"name":"foo"}' http://localhost:80/hello
		req := new(struct {
			Name string `json:"name"`
		})
		if err := c.Bind(req); err != nil {
			return err
		}
		// レスポンスの返却
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, " + req.Name + "!",
		})
	})

	// ミドルウェアの登録
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("sendai")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("go")) == 1 {
			// 認証成功
			return true, nil
		}
		// 認証失敗時はfalseを返す
		return false, nil
	}))

	// サーバーの起動
	e.Logger.Fatal(e.Start(":80"))
}
