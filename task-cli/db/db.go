package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(path string) (*taskDB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
	if err != nil {
		return nil, err
	}
	t := taskDB{db, path}
	if !t.tableExists("tasks") {
		err = t.createTable()
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func DeleteDB(path string) error {
	return os.Remove(filepath.Join(path, "tasks.db"))
}

func (t *taskDB) createTable() error {
	sqlStmt := `
		CREATE TABLE tasks (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"description" TEXT NOT NULL,
			"status" TEXT,
			"created" DATETIME DEFAULT CURRENT_TIMESTAMP,
			"updated" DATETIME
		)`

	_, err := t.DB.Exec(sqlStmt)
	return err
}

func (t *taskDB) CreateTask(name string) error {
	sqlStmt := `
		INSERT INTO tasks (description, status)
		VALUES ($1, $2)`

	_, err := t.DB.Exec(sqlStmt,
		name,
		todo.String(),
	)
	return err
}

func (t *taskDB) UpdateTask(id int, name string, statusIdx int) error {
	sqlStmt := `
		UPDATE tasks
		SET description = $1, status = $2, updated = CURRENT_TIMESTAMP
		WHERE id = $3`

	statusStr := mkStatus(statusIdx)
	_, err := t.DB.Exec(sqlStmt, name, statusStr, id)
	return err
}

func (t *taskDB) GetTasks() ([]task, error) {
	sqlStmt := `
		SELECT * FROM tasks`

	rows, err := t.DB.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	var tasks []task
	for rows.Next() {
		var t task
		err = rows.Scan(
			&t.ID,
			&t.Description,
			&t.Status,
			&t.Created,
			&t.Updated,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *taskDB) GetTask(id int) (task, error) {
	var task task
	sqlStmt := `
		SELECT * FROM tasks
		WHERE id = $1`

	err := t.DB.QueryRow(sqlStmt, id).
		Scan(
			&task.ID,
			&task.Description,
			&task.Status,
			&task.Created,
			&task.Updated,
		)
	return task, err
}

func (t *taskDB) DeleteTask(id int) error {
	sqlStmt := `
		DELETE FROM tasks
		WHERE id = $1`

	_, err := t.DB.Exec(sqlStmt, id)
	return err
}
