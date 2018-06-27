package main

import (
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"log"
	"time"
)

type Item struct {
	Room string
	Time string
}

func main() {
	f := firego.New("https://sao-unv.firebaseio.com", nil)
	firego.TimeoutDuration = time.Minute

	var v Item
	if err := f.Child("s1230008").Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", v)
}
