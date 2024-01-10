package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// TODOを作成する
// curl -X POST -H "Content-Type: application/json" -d '{"task":"task1"}' localhost/todo
func create(c echo.Context, db *sql.DB) error {
	// JSONをパースしてTODOを作成する
	req := new(todo)
	if err := c.Bind(req); err != nil {
		return err
	}
	// TODOのIDは自動採番される
	res, err := db.Exec("INSERT INTO todos (task) VALUES (?)", req.Task)
	// TODOの作成に失敗した場合はエラーを返す
	if err != nil {
		return err
	}
	// TODOの作成に成功した場合は作成したTODOを返す
	id, err := res.LastInsertId()
	return c.JSON(http.StatusOK, todo{ID: int(id), Task: req.Task})
}

// TODOを全件取得する
// curl localhost/todo
func readAll(c echo.Context, db *sql.DB) error {
	// TODOを全件取得する
	res := []todo{}
	rows, err := db.Query("SELECT id,task FROM todos order by id asc")
	if err != nil {
		return err
	}
	defer rows.Close()

	// TODOを1件ずつ取得する
	for rows.Next() {
		var id int
		var task string
		if err := rows.Scan(&id, &task); err != nil {
			return err
		}
		res = append(res, todo{ID: id, Task: task})
	}
	return c.JSON(http.StatusOK, res)
}

// TODOを1件取得する
// curl localhost/todo/1
func read(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	// TODOを1件取得する
	res := todo{}
	row := db.QueryRow("SELECT id,task FROM todos WHERE id = ?", id)
	if err := row.Scan(&res.ID, &res.Task); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// TODOを更新する
// curl -X PUT -H "Content-Type: application/json" -d '{"task":"task2"}' localhost/todo/1
func update(c echo.Context, db *sql.DB) error {
	// JSONをパースしてTODOを作成する
	req := new(todo)
	if err := c.Bind(req); err != nil {
		return err
	}
	id := c.Param("id")
	// TODOの更新
	_, err := db.Exec("UPDATE todos SET task = ? WHERE id = ?", req.Task, id)

	// TODOの更新に失敗した場合はエラーを返す
	if err != nil {
		return err
	}
	return nil
}

// TODOを削除する
// curl -X DELETE localhost/todo/1
func delete(c echo.Context, db *sql.DB) error {
	// TODOの削除
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)

	// TODOの削除に失敗した場合はエラーを返す
	if err != nil {
		return err
	}
	return nil
}
