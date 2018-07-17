package get

import (
	"fmt"
	"log"
	"time"

	"github.com/zabawaba99/firego"
)

type Command struct{}

func (f *Command) Help() string {
	return "Usage: whroom get <student_id>"
}

func (f *Command) Run(args []string) int {
	if len(args) < 1 {
		log.Println("Please input studnet id\nex. whroom get s1240215")
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

func (f *Command) Synopsis() string {
	return "Print the room where the student is in."
}
