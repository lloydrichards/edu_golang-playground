package db

import (
	"database/sql"
	"time"
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
	ID          uint
	Description string
	Status      string
	Created     time.Time
	Updated     *time.Time
}

type taskDB struct {
	db      *sql.DB
	dataDir string
}
