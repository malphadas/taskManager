/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"
	"strconv"
	"taskManager/config"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long: `Add a new task to the list by providing a task description`,
	Run: addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	
	file, err := os.OpenFile(config.TasksFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id, err := os.ReadFile(config.CounterFile)
	if err != nil {
		panic(err)
	}

	idInt, err := strconv.Atoi(string(id))
	if err != nil {
		panic(err)
	}

	idInt++
	os.WriteFile(config.CounterFile, []byte(strconv.Itoa(idInt)), 0644)

	newID := strconv.Itoa(idInt)
	task := args[0]
	writer.Write([]string{newID, task, time.Now().Format(time.RFC3339), "false"})


}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
