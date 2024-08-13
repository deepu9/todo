package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func DeleteTodoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [todoNames]",
		Short: "Delete a todo or list of todos.",
		Long:  "Delete a todo or list of todos. Provide multiple todos separated by space. This deletes the file with the given name.",
		Run:   deleteTodo,
	}
}

func deleteTodo(cmd *cobra.Command, args []string) {
	dataDirPath := utils.GetDataPath()

	for _, todoName := range args {
		todoFilePath := filepath.Join(
			dataDirPath,
			fmt.Sprintf("%s%s", todoName, utils.TodoFileExtension),
		)

		// Check whether the file exists
		if _, err := utils.CheckPathExists(todoFilePath); !err {
			fmt.Printf("%s doesn't exists.\n", todoFilePath)
			continue
		}

		if err := os.Remove(todoFilePath); err != nil {
			panic(err)
		}

		fmt.Printf("%s has been deleted.\n", todoFilePath)
	}
}
