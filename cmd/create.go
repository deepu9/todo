package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func CreateTodoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create [todoName]",
		Short: "Creates a new TODO",
		Long:  "Creates a new TODO. This will create a new file to store the todo items.",
		Args:  cobra.MinimumNArgs(1),
		Run:   createTodo,
	}
}

func createTodo(cmd *cobra.Command, args []string) {
	dataDirPath := utils.GetDataPath()

	for index, todoName := range args {
		todoFilePath := filepath.Join(
			dataDirPath,
			fmt.Sprintf("%s%s", todoName, utils.TodoFileExtension),
		)

		// Check whether the file exists
		if _, exists := utils.CheckPathExists(todoFilePath); exists {
			fmt.Printf("%d. %s exists.\n", index, todoFilePath)
			continue
		}

		file, err := os.Create(todoFilePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		fmt.Printf("%d. %s has been created.\n", index, todoFilePath)
	}
}
