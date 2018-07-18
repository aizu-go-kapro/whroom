package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aizu-go-kapro/whroom/get"
	"github.com/aizu-go-kapro/whroom/watch"
	"github.com/ktr0731/go-semver"
	"github.com/ktr0731/go-updater"
	"github.com/ktr0731/go-updater/brew"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
)

const version = "0.1.0"

func main() {
	ctx := context.Background()
	means, err := updater.NewMeans(brew.HomebrewMeans("aizu-go-kapro/whroom", "whroom"))
	if err != nil {
		fmt.Printf("could not create updater means. continue without update checking: %s\n", err)
	} else {
		upd := updater.New(semver.MustParse(version), means)
		updatable, latest, err := upd.Updatable(ctx)
		if err != nil {
			fmt.Printf("could not check whether can be updated. continue: %s\n", err)
		} else {
			if updatable {
				if err := upd.Update(ctx); err != nil {
					fmt.Println(errors.Wrap(err, "could not update"))
				}
				fmt.Printf("updating from %s to %s\n", version, latest)
			} else {
				fmt.Println("already latest")
			}
		}
	}

	c := cli.NewCLI("whroom", version)
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
