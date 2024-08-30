package db

import (
	"log"
	"os"

	gap "github.com/muesli/go-app-paths"
)

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

func (t *taskDB) tableExists(name string) bool {
	if _, err := t.DB.Query("SELECT * FROM tasks"); err == nil {
		return true
	}
	return false
}
