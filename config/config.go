package config

import (
	"os"
	"path/filepath"
)

var (
	DataDir     string
	TasksFile   string
	CounterFile string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DataDir = filepath.Join(home, "Projects", "taskManager")
	TasksFile = filepath.Join(DataDir, "tasks.csv")
	CounterFile = filepath.Join(DataDir, "id.counter")
}
