package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

func (s status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

type task struct {
	ID      uint
	Name    string
	Project string
	Status  string
	Created time.Time
}

type taskDB struct {
	db      *sql.DB
	dataDir string
}

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

func (t *taskDB) tableExists(name string) bool {
	if _, err := t.db.Query("SELECT * FROM tasks"); err == nil {
		return true
	}
	return false
}

func (t *taskDB) createTable() error {
	_, err := t.db.Exec(`CREATE TABLE "tasks" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`)
	return err
}

func (t *taskDB) insert(name, project string) error {
	// We don't care about the returned values, so we're using Exec. If we
	// wanted to reuse these statements, it would be more efficient to use
	// prepared statements. Learn more:
	// https://go.dev/doc/database/prepared-statements
	_, err := t.db.Exec(
		"INSERT INTO tasks(name, project, status, created) VALUES( ?, ?, ?, ?)",
		name,
		project,
		todo.String(),
		time.Now())
	return err
}

func (t *taskDB) getTasks() ([]task, error) {
	return nil, nil
}

func (t *taskDB) getTask(id uint) (task, error) {
	var task task
	err := t.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).
		Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
		)
	return task, err
}

func initTaskDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}

func SetupPath() string {
	scope := gap.NewScope(gap.User, "tasks")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}
	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		taskDir, _ = os.UserHomeDir()
	}

	err = initTaskDir(taskDir)
	if err != nil {
		log.Fatal(err)
	}
	return taskDir
}
