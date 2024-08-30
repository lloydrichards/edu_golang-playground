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

var mockTasks = []Task{
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
		want Task
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
		t.Run("get task: "+tc.want.Description, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			_, err := mockDB.CreateTask(tc.want.Description)
			if err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			task, err := mockDB.GetTask(tc.want.ID)
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
		tasks, err := mockDB.GetTasks()
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
		want Task
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
		t.Run("delete task: "+tc.want.Description, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			err := mockDB.DeleteTask(tc.want.ID)
			if err != nil {
				t.Fatalf("we ran into an unexpected error: %v", err)
			}
			_, err = mockDB.GetTask(tc.want.ID)
			if err == nil {
				t.Fatalf("task was not deleted")
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		before Task
		after  Task
	}{
		{
			before: mockTasks[0],
			after: Task{
				ID:          mockTasks[0].ID,
				Description: "new description",
				Status:      mockTasks[0].Status,
			},
		},
		{
			before: mockTasks[1],
			after: Task{
				ID:          mockTasks[1].ID,
				Description: mockTasks[1].Description,
				Status:      inProgress.String(),
			},
		},
		{
			before: mockTasks[2],
			after: Task{
				ID:          mockTasks[2].ID,
				Description: mockTasks[2].Description,
				Status:      done.String(),
			},
		},
	}

	for _, tc := range tests {

		t.Run("update task: "+tc.after.Description, func(t *testing.T) {
			mockDB := setup()
			defer teardown(mockDB)
			status, err := StatusFromString(tc.after.Status)
			if err != nil {
				t.Fatalf("The status is invalid: %v", err)
			}

			err = mockDB.UpdateTask(tc.after.ID, tc.after.Description, status)
			if err != nil {
				t.Fatalf("Failed to update task: %v", err)
			}
			task, err := mockDB.GetTask(tc.after.ID)
			if err != nil {
				t.Fatalf("Failed to get task: %v", err)
			}
			if task.Description != tc.after.Description {
				t.Fatalf("got: %s, want: %s", task.Description, tc.after.Description)
			}
			if task.Status != tc.after.Status {
				t.Fatalf("got: %s, want: %s", task.Status, tc.after.Status)
			}
			if task.Updated == nil {
				t.Fatalf("updated time was not set")
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

	for _, task := range mockTasks {
		_, err := t.CreateTask(task.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &t
}

func teardown(tDB *taskDB) {
	tDB.DB.Close()
	os.Remove(tDB.DataDir)
}
