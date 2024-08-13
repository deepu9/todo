package item

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func ItemAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add [itemNames]",
		Short: "Add one or more new items to a TODO",
		Long:  "Add one or more new items to the given TODO. Append the item to the TODO file.",
		Args:  cobra.MinimumNArgs(1),
		Run:   addItem,
	}
}

func addItem(cmd *cobra.Command, args []string) {
	todoName := cmd.Flags().Lookup(utils.TodoNameLabel).Value

	// Check whether a todo file exists
	dataDirPath := utils.GetDataPath()
	todoFilePath := filepath.Join(
		dataDirPath,
		fmt.Sprintf("%s%s", todoName, utils.TodoFileExtension),
	)
	
	// open the todo file
	fileHandler, err := os.OpenFile(todoFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer fileHandler.Close()

	fileWriter := bufio.NewWriter(fileHandler)

	for _, line := range args {
		fileWriter.WriteString(line)
		fileWriter.WriteString("\n")
		fileWriter.Flush()
	}

	fmt.Printf("Successfully written to file: %s\n", todoFilePath)
}
