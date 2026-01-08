/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"os"
	"taskManager/cmd"
	"taskManager/config"
)

func main() {
	// Check if file exists and get its size
	fileInfo, err := os.Stat(config.TasksFile)
	fileExists := err == nil
	isEmpty := !fileExists || fileInfo.Size() == 0

	file, err := os.OpenFile(config.TasksFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Only write header if file is empty or doesn't exist
	if isEmpty {
		writer.Write([]string{"ID", "Task", "Created", "Completed"})
	}

	cmd.Execute()
}
