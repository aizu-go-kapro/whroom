package watch

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

type Command struct {
	FirebaseURL   string
	StudentID     string
	WifiInterface string
	Duration      time.Duration
}

func (c *Command) Run(args []string) int {
	for range time.Tick(c.Duration) {
		room, ok := getRoomFromAPs(listAP(c.WifiInterface))
		if !ok {
			fmt.Fprintln(os.Stderr, "could not estimate which room you are in")
			continue
		}
		if err := Save(c.FirebaseURL, c.StudentID, room, time.Now()); err != nil {
			fmt.Fprintln(os.Stderr, errors.Wrap(err, "could not save the room info"))
		}
	}

	return 0
}

func (c *Command) Help() string {
	return `Usage: whroom watch

Note: Usually you don't need to run this command from shell. To start location
logging, ref docs.` // TODO: add ref.
}

func (c *Command) Synopsis() string {
	return "Watch the Wi-Fi connection and send to server periodically."
}

var bssids = map[Room][]string{
	M7:        {"6c:dd:30:37:39:ab", "6c:dd:30:37:39:af", "6c:dd:30:37:39:a2"},
	S1:        {"00:a7:42:ad:c3:6b", "00:a7:42:ad:c3:6f"},
	S2:        {"6c:dd:30:49:6b:44", "6c:dd:30:49:6b:4f"},
	S3:        {"4c:77:6d:17:76:24", "4c:77:6d:17:76:20"},
	S4:        {"4c:77:6d:17:76:2b", "4c:77:6d:17:76:2f"},
	S5:        {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S6:        {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S7:        {"4c:77:6d:17:66:c4", "4c:77:6d:17:66:c0"},
	S8:        {"4c:77:6d:44:e4:0", "4c:77:6d:44:e4:00"},
	CALL1:     {"f8:b7:e2:d3:fb:4b", "f8:b7:e2:d3:fb:4f"},
	CALL2:     {"f8:b7:e2:d3:ce:c4", "f8:b7:e2:d3:ce:c0"},
	ILab1:     {"6c:dd:30:49:61:44", "6c:dd:30:49:61:40"},
	ILab2:     {"f8:b7:e2:d3:fb:a4", "f8:b7:e2:d3:fb:a0"},
	Std1:      {"f8:b7:e2:d3:d3:c4", "f8:b7:e2:d3:d3:c0"},
	Std2:      {"6c:dd:30:49:61:04", "6c:dd:30:49:61:00"},
	Std5:      {"6c:dd:30:49:59:44", "6c:dd:30:49:59:40"},
	Std6:      {"6c:dd:30:49:38:64", "6c:dd:30:49:38:60"},
	Cafeteria: {"f8:b7:e2:cc:7a:c4", "f8:b7:e2:cc:7a:c0"},
}

func getRoomFromAPs(aps []*AP) (Room, bool) {
	sort.Slice(aps, func(i, j int) bool {
		return aps[i].Signal < aps[j].Signal
	})
	pp.Print(aps)

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
