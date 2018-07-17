package main

import "fmt"

func main() {
	d := make(chan struct{})
	c := make(chan map[string]map[string]interface{})
	go func() {
		for v := range c {
			fmt.Println(v)
		}
		close(d)
	}()
	iw(c)
	<-d
}
