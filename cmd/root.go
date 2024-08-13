package cmd

import (
	"fmt"
	"os"

	"github.com/deepu9/todo/cmd/item"
	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A TODO command line tool.",
	Long:  "A TODO command line tool to create, add, delete and list the TODO's. The TODO's get stored in a file.",
}

var todoName string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&todoName, utils.TodoNameLabel, "n", "", "Provide todo name to access it's items.")

	// Todo commands
	rootCmd.AddCommand(CreateTodoCommand())
	rootCmd.AddCommand(ListTodoCommand())
	rootCmd.AddCommand(DeleteTodoCommand())

	rootCmd.AddCommand(item.ItemMainCommand())
}
