package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetTask(t *testing.T) {
	tests := []struct {
		want task
	}{
		{
			want: task{
				ID:      1,
				Name:    "get milk",
				Project: "groceries",
				Status:  todo.String(),
			},
		},
		{
			want: task{
				ID:      1,
				Name:    "get eggs",
				Project: "groceries",
				Status:  todo.String(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.want.Name, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			if err := mockDB.insert(tc.want.Name, tc.want.Project); err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			task, err := mockDB.getTask(tc.want.ID)
			fmt.Print(task)
			if err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			tc.want.Created = task.Created
			if !reflect.DeepEqual(task, tc.want) {
				t.Fatalf("got: %#v, want: %#v", task, tc.want)
			}
		})
	}

}

func setup() *taskDB {
	path := filepath.Join(os.TempDir(), "test.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	t := taskDB{db, path}
	if !t.tableExists("tasks") {
		err := t.createTable()
		if err != nil {
			log.Fatal(err)
		}
	}
	return &t
}

func teardown(tDB *taskDB) {
	tDB.db.Close()
	os.Remove(tDB.dataDir)
}
