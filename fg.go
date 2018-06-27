package main

import (
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	"time"
)

func main() {
	k := "s1230008"
	f := firego.New("https://sao-unv.firebaseio.com", nil)
	firego.TimeoutDuration = time.Minute

	var v map[string]interface{}
	if err := f.Child(k).Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", v["room"])
	fmt.Printf("%s\n", v["timestamp"])
}
