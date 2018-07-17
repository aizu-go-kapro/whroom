package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"log"
)

var minimumThreshold = -100
var wifiInterface = "wlp2s0"

func iw(out chan map[string]map[string]interface{}) {
	s, _ := RunCommand(10*time.Second, "/sbin/iw dev "+wifiInterface+" scan -u")
	name := ""
	signal := 0
	datas := make(map[string]map[string]interface{})
	datas["wifi"] = make(map[string]interface{})
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "(on") {
			name = strings.Split(strings.Split(line, "(")[0], "BSS")[1]
			name = strings.ToLower(name)
			name = strings.TrimSpace(name)
		} else if strings.Contains(line, "signal:") {
			foo := strings.Split(line, "signal:")[1]
			foo = strings.Split(foo, ".")[0]
			foo = strings.TrimSpace(foo)
			var err error
			signal, err = strconv.Atoi(foo)
			if err != nil {
				panic(err)
			}
		}
		if name != "" && signal != 0 {
			if signal < minimumThreshold {
				continue
			}
			datas["wifi"][name] = signal
		}
	}
	out <- datas
}

func RunCommand(tDuration time.Duration, commands string) (string, string) {
	log.Print(commands)
	command := strings.Fields(commands)
	cmd := exec.Command(command[0])
	if len(command) > 0 {
		cmd = exec.Command(command[0], command[1:]...)
	}
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(tDuration):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
		}
		log.Printf("%s killed as timeout reached", commands)
	case err := <-done:
		if err != nil {
			log.Printf("%s: %s", err.Error(), commands)
		} else {
			log.Printf("%s done gracefully without error", commands)
		}
	}
	return strings.TrimSpace(outb.String()), strings.TrimSpace(errb.String())
}
