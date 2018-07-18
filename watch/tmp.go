package watch

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

type AP struct {
	BSSID  string
	Signal int
}

var cmdstub = ""

func RunCommand(tDuration time.Duration, commands string) (string, string) {
	if cmdstub != "" {
		return cmdstub, ""
	}

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
