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
	"github.com/ktr0731/go-updater/github"
	"github.com/mitchellh/cli"
	"github.com/pkg/errors"
)

const version = "0.1.1"

func main() {
	ctx := context.Background()
	means, err := updater.SelectAvailableMeansFrom(
		ctx,
		brew.HomebrewMeans("aizu-go-kapro/whroom", "whroom"),
		github.GitHubReleaseMeans("aizu-go-kapro", "whroom", github.TarGZIPDecompresser),
	)
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

	config, args, err := getConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	c := cli.NewCLI("whroom", version)
	c.Args = args
	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return &get.Command{
				FirebaseURL: config.FirebaseURL,
			}, nil
		},
		"watch": func() (cli.Command, error) {
			return &watch.Command{
				FirebaseURL:   config.FirebaseURL,
				StudentID:     config.StudentID,
				WifiInterface: config.WifiInterface,
				Duration:      config.Duration,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
