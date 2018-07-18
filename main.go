package main

import (
	"log"
	"os"

	"github.com/aizu-go-kapro/whroom/get"
	"github.com/aizu-go-kapro/whroom/watch"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("whroom", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return &get.Command{}, nil
		},
		"watch": func() (cli.Command, error) {
			return &watch.Command{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
