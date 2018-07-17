package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/cli"
	"gopkg.in/zabawaba99/firego.v1"
)

type StudentCommand struct{}

func (f *StudentCommand) Help() string {
	return "Usage: whroom student <student_id>"
}

func (f *StudentCommand) Run(args []string) int {
	if len(args) < 1 {
		log.Println("Please input studnet number\nex. whroom student s1240215")
		return 1
	}
	k := args[0]
	fg := firego.New("https://sao-unv.firebaseio.com", nil)
	firego.TimeoutDuration = time.Minute

	var v map[string]interface{}
	if err := fg.Child(k).Value(&v); err != nil {
		log.Println(err)
		return 1
	}
	if v == nil {
		log.Println("the student not record any location log yet.")
		return 1
	}

	fmt.Printf("room: %s\n", v["room"])
	fmt.Printf("timestamp: %s\n", v["timestamp"])

	return 0
}

func (f *StudentCommand) Synopsis() string {
	return "Print the room where the student is in."
}

func main() {
	c := cli.NewCLI("whroom", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"student": func() (cli.Command, error) {
			return &StudentCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
