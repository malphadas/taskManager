/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"taskManager/config"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List incomplete tasks from the database, add flag --all (-a) to list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		if all {
			listAllTasks(cmd, args)
		} else {
			listIncompleteTasks(cmd, args)
		}
	},
}

func listIncompleteTasks(cmd *cobra.Command, args []string) {
	in, err := os.Open(config.TasksFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	reader := csv.NewReader(in)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()

	for i, record := range records {


		if i == 0 {
			fmt.Fprintln(writer, strings.Join(record, " \t"))
			continue
		}

		ts, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			panic(err)
		}

		record[2] = timediff.TimeDiff(ts)
		if record[3] == "false" {
			fmt.Fprintln(writer, strings.Join(record, " \t"))
		}
	}
}

func listAllTasks(cmd *cobra.Command, args []string) {

	in, err := os.Open(config.TasksFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	reader := csv.NewReader(in)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()

	for i, record := range records {


		if i == 0 {
			fmt.Fprintln(writer, strings.Join(record, " \t"))
			continue
		}

		ts, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			panic(err)
		}

		record[2] = timediff.TimeDiff(ts)
		fmt.Fprintln(writer, strings.Join(record, " \t"))
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "List all tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
