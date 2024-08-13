package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func ListTodoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all the todo's",
		Long:  "List all the todo's. Get all the files and print the file names.",
		Run:   listTodos,
	}
}

func listTodos(cmd *cobra.Command, args []string) {
	// Get data directory
	dataDirPath := utils.GetDataPath()

	// Check whether there are any files in the directory
	files, err := os.ReadDir(dataDirPath)
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Println("no todos found.")
		os.Exit(1)
	}

	fmt.Println("List of TODOS:")

	// Read all the files from the directory
	for index, file := range files {
		fileName := file.Name()
		extension := filepath.Ext(fileName)

		fmt.Printf("%d. %s\n", index, strings.TrimSuffix(fileName, extension))
	}
}
