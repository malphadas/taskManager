/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"
	"taskManager/config"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as completed",
	Long: `Mark a task as completed by providing the task ID`,
	Run: completeTask,
}

func completeTask(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		panic("missing task id")
	}
	targetID := args[0]

	in, err := os.Open(config.TasksFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	tmpPath := config.TasksFile + ".tmp"
	out, err := os.Create(tmpPath) // truncates/creates temp
	if err != nil {
		panic(err)
	}
	defer out.Close()

	reader := csv.NewReader(in)
	writer := csv.NewWriter(out)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, record := range records {
		// keep header as-is
		if i == 0 {
			if err := writer.Write(record); err != nil {
				panic(err)
			}
			continue
		}

		if len(record) >= 4 && record[0] == targetID {
			record[3] = "true"
		}

		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
	if err := out.Sync(); err != nil {
		panic(err)
	}
	if err := out.Close(); err != nil {
		panic(err)
	}
	if err := in.Close(); err != nil {
		panic(err)
	}

	if err := os.Rename(tmpPath, config.TasksFile); err != nil {
		panic(err)
	}
}


func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
