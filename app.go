package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "todo",
		Usage: "manage my own tasks",
		Action: func(c *cli.Context) error {
			fmt.Println("Dariush's own todo manager")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add your todo",
				Action:  Add,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "tasks list",
				Action:  List,
			},
			{
				Name:   "doing",
				Usage:  "change status of task to doing",
				Action: Doing,
			},
			{
				Name:   "done",
				Usage:  "change status of task to done",
				Action: Done,
			},
			{
				Name:   "archive",
				Usage:  "change status of task to archived",
				Action: Archive,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
