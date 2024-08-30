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

var mockTasks = []task{
	{
		ID:          1,
		Description: "get milk",
		Status:      todo.String(),
	},
	{
		ID:          2,
		Description: "get eggs",
		Status:      todo.String(),
	},
	{
		ID:          3,
		Description: "get bread",
		Status:      todo.String(),
	},
}

func TestGetTask(t *testing.T) {
	tests := []struct {
		want task
	}{
		{
			want: mockTasks[0],
		},
		{
			want: mockTasks[1],
		},
		{
			want: mockTasks[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.want.Description, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			if err := mockDB.insert(tc.want.Description); err != nil {
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

func TestGetTasks(t *testing.T) {

	t.Run("get all tasks", func(t *testing.T) {
		mockDB := setup()
		defer teardown(mockDB)
		tasks, err := mockDB.getTasks()
		if err != nil {
			t.Fatalf("we ran into an unexpected error: %v", err)
		}
		if len(tasks) != len(mockTasks) {
			t.Fatalf("got: %d, want: %d", len(tasks), len(mockTasks))
		}
	},
	)
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		want task
	}{
		{
			want: mockTasks[0],
		},
		{
			want: mockTasks[1],
		},
		{
			want: mockTasks[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.want.Description, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			if err := mockDB.insert(tc.want.Description); err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			err := mockDB.deleteTask(tc.want.ID)
			if err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			_, err = mockDB.getTask(tc.want.ID)
			if err == nil {
				t.Fatalf("task was not deleted")
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
	log.Print("checking if table exists")
	if !t.tableExists("tasks") {
		log.Print("creating")
		err := t.createTable()
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("seeding db")
	for _, task := range mockTasks {
		err := t.insert(task.Description)
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
