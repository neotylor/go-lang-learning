/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
/* var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
} */

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		todos, err := repo.List()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(todos) == 0 {
			fmt.Println("No todos found.")
			return
		}
		for _, t := range todos {
			status := "[ ]"
			if t.Completed {
				status = "[x]"
			}
			fmt.Printf("%s %d: %s\n", status, t.ID, t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
