package main

import "github.com/deepu9/todo/cmd"

func main() {
	// DONE - todo create {todoName} - This will create a file with the given name
	// DONE - todo -n {todoName} item add {itemName} - This will add a new item to the given todo
	// DONE - todo -n {todoName} item list - Prints all the items of a given todo
	// DONE - todo list - Prints all the todo names(parents)
	// DONE - todo delete {todoName} - Remove a todo file.
	cmd.Execute()
}
