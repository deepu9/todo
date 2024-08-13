package item

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func ItemListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all the items of a TODO",
		Long:  "List all the items of a TODO. Display the list of items of a TODO",
		Run:   listItems,
	}
}

func listItems(cmd *cobra.Command, args []string) {
	todoName := cmd.Flags().Lookup(utils.TodoNameLabel).Value
	dataDirPath := utils.GetDataPath()
	todoFilePath := filepath.Join(
		dataDirPath,
		fmt.Sprintf("%s%s", todoName, utils.TodoFileExtension),
	)

	fileInfo, _ := os.Stat(todoFilePath)
	// Check whether the todo file is empty
	if fileInfo.Size() == 0 {
		fmt.Printf("no items found in %s.\n", todoName)
		os.Exit(1)
	}

	// Display the list of todo items
	file, err := os.Open(todoFilePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileReader := bufio.NewReader(file)
	index := 0

	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Printf("%d. %s", index, line)
		index++
	}
}
