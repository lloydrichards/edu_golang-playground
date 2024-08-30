package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Status int

const (
	todo Status = iota
	inProgress
	done
)

func (s Status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

func (s Status) Int() int {
	return int(s)
}

func StatusFromString(str string) (int, error) {
	switch str {
	case "todo":
		return todo.Int(), nil
	case "in progress":
		return inProgress.Int(), nil
	case "done":
		return done.Int(), nil
	default:
		return -1, fmt.Errorf("invalid status: %s", str)
	}
}

func mkStatus(i int) string {
	return [...]string{"todo", "in progress", "done"}[i]
}

type task struct {
	ID          int
	Description string
	Status      string
	Created     time.Time
	Updated     *time.Time
}

type taskDB struct {
	DB      *sql.DB
	DataDir string
}
