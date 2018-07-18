package watch

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/zchee/go-xdgbasedir"
	"github.com/zchee/go-xdgbasedir/home"
)

type Command struct{}

func (c *Command) Run(args []string) int {
	path, ok := getConfigFilePath()
	if !ok {
		fmt.Fprintln(os.Stderr, "could not find config file.") // TODO: add reference
		return 1
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "could not open %s", path))
		return 1
	}

	config, err := decodeConfig(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "failed to parse the config file"))
		return 1
	}

	for {
		room, ok := getRoomFromAPs(listAP(config.WifiInterface))
		if !ok {
			fmt.Fprintln(os.Stderr, "could not estimate which room you are in")
		}
		if err := Save(config.FirebaseURL, config.StudentID, room, time.Now()); err != nil {
			fmt.Fprintln(os.Stderr, errors.Wrap(err, "could not save the room info"))
		}
		<-time.After(config.Duration)
	}
}

func (c *Command) Help() string {
	return `Usage: whroom watch

Note: Usually you don't need to run this command from shell. To start location
logging, ref docs.` // TODO: add ref.
}

func (c *Command) Synopsis() string {
	return "Watch the Wi-Fi connection and send to server periodically."
}

func getConfigFilePath() (string, bool) {
	paths := []string{
		filepath.Join(xdgbasedir.ConfigHome(), "whroom", "config.toml"),
		filepath.Join(home.Dir(), ".whroom.toml"),
	}
	fmt.Println(paths)
	for _, path := range paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path, true
		}
	}
	return "", false
}

var bssids = map[Room][]string{
	M7:    {"6c:dd:30:37:39:ab", "6c:dd:30:37:39:af"},
	S1:    {"00:a7:42:ad:c3:6b", "00:a7:42:ad:c3:6f"},
	S2:    {"6c:dd:30:49:6b:44", "6c:dd:30:49:6b:4f"},
	S3:    {"4c:77:6d:17:76:24", "4c:77:6d:17:76:20"},
	S4:    {"4c:77:6d:17:76:2b", "4c:77:6d:17:76:2f"},
	S5:    {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S6:    {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S7:    {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S8:    {"4c:77:6d:44:e4:0", "4c:77:6d:44:e4:00"},
	CALL1: {"f8:b7:e2:d3:fb:4b", "f8:b7:e2:d3:fb:4f"},
	CALL2: {"f8:b7:e2:d3:ce:c4", "f8:b7:e2:d3:ce:c0"},
	iLab1: {"6c:dd:30:49:61:44", "6c:dd:30:49:61:40"},
	iLab2: {"f8:b7:e2:d3:fb:a4", "f8:b7:e2:d3:fb:a0"},
	std1:  {"f8:b7:e2:d3:d3:c4", "f8:b7:e2:d3:d3:c0"},
	std2:  {"6c:dd:30:49:61:04", "6c:dd:30:49:61:00"},
	std5:  {"6c:dd:30:49:59:44", "6c:dd:30:49:59:40"},
	std6:  {"6c:dd:30:49:38:64", "6c:dd:30:49:38:60"},
}

func getRoomFromAPs(aps []*AP) (Room, bool) {
	sort.Slice(aps, func(i, j int) bool {
		return aps[i].Signal < aps[j].Signal
	})

	var room Room
	var ok bool
	for _, ap := range aps {
		room, ok = getRoom(ap.BSSID)
		if ok {
			break
		}
	}

	return room, ok
}

func getRoom(bssid string) (Room, bool) {
	for room, bs := range bssids {
		for _, b := range bs {
			if bssid == b {
				return room, true
			}
		}
	}
	return Room(0), false
}
