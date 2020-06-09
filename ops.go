package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/dariubs/s2i"
)

// Add a todo item
func Add(c *cli.Context) error {
	td := Todo{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("todo: ")
	td.Task, _ = reader.ReadString('\n')
	td.Task = strings.Replace(td.Task, "\n", "", -1)
	td.Status = "active"
	td.Add()
	return nil
}

// List of todo items
func List(c *cli.Context) error {
	tds := []Todo{}
	tds, err = Get()
	if err != nil {
		return err
	}
	fmt.Printf("ID\tTask\t\tStatus\n")
	fmt.Printf("---\t---\t\t---\n")
	for _, v := range tds {
		fmt.Printf("%d\t%s\t\t%s\n", v.ID, v.Task, v.Status)
	}
	return nil
}

// Doing a task
func Doing(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	td.ChangeStatus("doing")
	return nil
}

// Done a task
func Done(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	td.ChangeStatus("done")
	return nil
}

// Archive a task
func Archive(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	td.ChangeStatus("archived")
	return nil
}
