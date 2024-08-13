## Todo Cli
This CLI tool will let you create multiple todos and each todo can have one or many items. If you want to use multiple words either for naming the todo or the item, then you need to use the double quotes around the todo name or item.

### Setup
Clone the repo and run `go build`. This will create an executable file with name **todo**.

### Modules used
1. [Cobra Cli](https://cobra.dev/)

### Usage
#### Todo files
1. Create one or more todo's with the following command
    ```bash
    ./todo create [todoName]
    ```
    ex.: `./todo create todo1 "new todo"` - Multiple todos

    ex.: `./todo create todo2` - Single todo
    
    This will create **data** folder in the root and create text files with the given names.
2. To list the todo's, run the following command:
    ```bash
    ./todo list
    ```
    This will display list of all the todos created under **data** folder. Display all the text file names.
3. You can delete a todo using the following command:
    ```bash
    ./todo delete [todoName]
    ```

    ex.: `./todo delete todo1`

    This will delete the entire text file from the **data** folder.

#### Todo Items

**NOTE:** To work with todo items, you need to know the todo/file name. All the item commands need the todo name.

1. To add individual items to a todo, run the following command:
    ```bash
    ./todo -n [todoName] item add [itemName]
    ```

    ex.: `./todo -n todo1 item add "This is a test" TestItem`

    This will add the given items into the todo1.txt
2. To view the items of a todo, use the following command:
    ```bash
    ./todo -n [todoName] item list
    ```

    ex.: `./todo -n todo1 item list`
