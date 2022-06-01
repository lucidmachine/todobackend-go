package main

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dropTableSql = `
	DROP TABLE IF EXISTS todos;
	`

	createTableSql = `
	CREATE TABLE IF NOT EXISTS todos (
		id 				TEXT		NOT NULL 	PRIMARY KEY,
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

	selectTodoByIdSql = `
	SELECT id, title, completed, "order", url
	FROM todos
	WHERE id = ?;
	`

	deleteTodosSql = `
	DELETE FROM todos;
	`
)

type Repository struct {
	db                 *sql.DB
	dropTableStmt      *sql.Stmt
	createTableStmt    *sql.Stmt
	insertTodoStmt     *sql.Stmt
	selectTodosStmt    *sql.Stmt
	selectTodoByIdStmt *sql.Stmt
	deleteTodoStmt     *sql.Stmt
}

func NewRepository(db *sql.DB) (*Repository, error) {
	dropTableStmt, err := db.Prepare(dropTableSql)
	if err != nil {
		return nil, err
	}

	createTableStmt, err := db.Prepare(createTableSql)
	if err != nil {
		return nil, err
	}

	insertTodoStmt, err := db.Prepare(insertTodoSql)
	if err != nil {
		return nil, err
	}

	selectTodosStmt, err := db.Prepare(selectTodosSql)
	if err != nil {
		return nil, err
	}

	selectTodoByIdStmt, err := db.Prepare(selectTodoByIdSql)
	if err != nil {
		return nil, err
	}

	deleteTodoStmt, err := db.Prepare(deleteTodosSql)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                 db,
		dropTableStmt:      dropTableStmt,
		createTableStmt:    createTableStmt,
		insertTodoStmt:     insertTodoStmt,
		selectTodosStmt:    selectTodosStmt,
		selectTodoByIdStmt: selectTodoByIdStmt,
		deleteTodoStmt:     deleteTodoStmt,
	}, nil
}

func (repo Repository) DropTable() error {
	_, err := repo.dropTableStmt.Exec()
	return err
}

func (repo Repository) CreateTable() error {
	_, err := repo.createTableStmt.Exec()
	return err
}

func (repo Repository) CreateTodo(todo Todo) error {
	_, err := repo.insertTodoStmt.Exec(todo.Id, todo.Title, todo.Completed, todo.Order, todo.Url)
	return err
}

func (repo Repository) GetTodos() ([]Todo, error) {
	rows, err := repo.selectTodosStmt.Query()
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

func (repo Repository) GetTodo(id uuid.UUID) (Todo, error) {
	row := repo.selectTodoByIdStmt.QueryRow(id)

	todo := Todo{}
	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.Order, &todo.Url)
	return todo, err
}

func (repo Repository) DeleteTodos() (int, error) {
	res, err := repo.deleteTodoStmt.Exec()
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
