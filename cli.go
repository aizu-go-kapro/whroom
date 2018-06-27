package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	"os"
	"time"
)

type Student struct{}

func (f *Student) Help() string {
	return "whroom student"
}

func (f *Student) Run(args []string) int {
	if len(args) < 1 {
		fmt.Printf("Please studnet number\n")
		return 0
	}
	k := args[0]
	fg := firego.New("https://sao-unv.firebaseio.com", nil)
	firego.TimeoutDuration = time.Minute

	var v map[string]interface{}
	if err := fg.Child(k).Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("room: %s\n", v["room"])
	fmt.Printf("timestamp: %s\n", v["timestamp"])
	return 0
}

func (f *Student) Synopsis() string {
	return "Print \"room and timestamp \""
}

func main() {
	c := cli.NewCLI("whroom", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"student": func() (cli.Command, error) {
			return &Student{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
