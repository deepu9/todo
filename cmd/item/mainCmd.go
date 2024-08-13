package item

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepu9/todo/utils"
	"github.com/spf13/cobra"
)

func ItemMainCommand() *cobra.Command {
	itemMainCommand := &cobra.Command{
		Use:              "item",
		Short:            "Subcommand to access todo items.",
		Long:             "Subcommand to access todo items.",
		PersistentPreRun: validateForTodoName,
		Run:              func(cmd *cobra.Command, args []string) {},
	}

	itemMainCommand.AddCommand(ItemAddCommand())
	itemMainCommand.AddCommand(ItemListCommand())

	return itemMainCommand
}

func validateForTodoName(cmd *cobra.Command, args []string) {
	todoNameFlag := cmd.Flags().Lookup(utils.TodoNameLabel)
	if todoNameFlag == nil {
		fmt.Println("Missing todo name")
		os.Exit(1)
	}

	todoName := todoNameFlag.Value.String()
	if todoName == "" {
		fmt.Println("Missing todo name")
		os.Exit(1)
	}

	dataDirPath := utils.GetDataPath()
	todoFilePath := filepath.Join(
		dataDirPath,
		fmt.Sprintf("%s%s", todoName, utils.TodoFileExtension),
	)

	// Check whether the todo file exists
	_, exists := utils.CheckPathExists(todoFilePath)
	if !exists {
		fmt.Printf("%s doesn't exists.\n", todoFilePath)
		os.Exit(1)
	}
}
