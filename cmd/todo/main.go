package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/buguno/todo-cli/internal/todo"
)

const todosFile = ".todos.json"

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todosFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")

		err := todos.Store(todosFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
