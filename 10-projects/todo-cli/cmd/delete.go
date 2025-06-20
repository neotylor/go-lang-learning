/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
/* var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
} */

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := todo.NewRepository()
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		err = repo.Delete(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Todo deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
