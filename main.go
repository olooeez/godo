package main

import (
	"log"
	"os"

	"github.com/olooeez/godo/internal/cli"
)

func main() {
	tasksFilePath := os.Getenv("HOME") + "/.godo.csv"

	app := cli.TaskManagerCli(tasksFilePath)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
