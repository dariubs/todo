package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dariubs/s2i"
	"github.com/jedib0t/go-pretty/table"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

// Add a todo item
func Add(c *cli.Context) error {
	td := Todo{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("todo: ")
	td.Task, _ = reader.ReadString('\n')
	td.Task = strings.Replace(td.Task, "\n", "", -1)
	td.Status = "active"

	prompt := promptui.Select{
		Label: "Select Status",
		Items: []string{"active", "doing", "done", "archived"},
	}

	_, td.Status, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	fmt.Printf("Task %s with status %q added :)\n", td.Task, td.Status)
	td.Add()

	return nil
}

// List of todo items
func List(c *cli.Context) error {
	tds := []Todo{}
	tds, err = Get(c.Args().Get(0))
	if err != nil {
		return err
	}
	data := []table.Row{}
	for _, v := range tds {
		x := table.Row{v.ID, v.Task, v.Status}
		data = append(data, x)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Task", "Status"})
	t.AppendRows(data)
	t.Render()
	return nil
}

// Change status of a task
func Change(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	prompt := promptui.Select{
		Label: "Select Status",
		Items: []string{"active", "doing", "done", "archived"},
	}

	_, td.Status, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}
	err = td.ChangeStatus(td.Status)
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d changed to %s\n", td.ID, td.Status)
	return nil
}

// Doing a task
func Doing(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	err := td.ChangeStatus("doing")
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d changed to doing\n", td.ID)
	return nil
}

// Done a task
func Done(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	err := td.ChangeStatus("done")
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d changed to done\n", td.ID)
	return nil
}

// Archive a task
func Archive(c *cli.Context) error {
	td := Todo{ID: s2i.ParseUint(c.Args().Get(0), 0)}
	err := td.ChangeStatus("archived")
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d archived\n", td.ID)
	return nil
}
