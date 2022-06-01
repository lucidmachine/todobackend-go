package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dropTableSql = `
	DROP TABLE IF EXISTS todos;
	`

	createTableSql = `
	CREATE TABLE IF NOT EXISTS todos (
		id 				INTEGER	NOT NULL 	PRIMARY KEY,
		title 		TEXT 		NOT NULL,
		completed INTEGER NOT NULL,
		"order"		INTEGER NOT NULL,
		url 			TEXT 		NOT NULL
	);`

	insertTodoSql = `
	INSERT INTO todos
	(id, title, completed, "order", url)
	VALUES
	(?, ?, ?, ?, ?);
	`

	selectTodosSql = `
	SELECT id, title, completed, "order", url
	FROM todos;
	`
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (repo Repository) DropTable() error {
	_, err := repo.db.Exec(dropTableSql)
	return err
}

func (repo Repository) CreateTable() error {
	_, err := repo.db.Exec(createTableSql)
	return err
}

func (repo Repository) CreateTodo(todo Todo) (int, error) {
	res, err := repo.db.Exec(insertTodoSql, nil, todo.Title, todo.Completed, todo.Order, todo.Url)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (repo Repository) GetTodos() ([]Todo, error) {
	rows, err := repo.db.Query(selectTodosSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.Order, &todo.Url)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
