package watch

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ssidLineRegexp = regexp.MustCompile(`^.*\s+((\w\w:){5}\w\w)\s+(-\d+).*$`)

func listAP(ifname string) []*AP {
	var aps []*AP

	stdout, _ := RunCommand(10*time.Second, "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport --scan")
	scn := bufio.NewScanner(strings.NewReader(stdout))
	for scn.Scan() {
		matches := ssidLineRegexp.FindStringSubmatch(scn.Text())
		if matches == nil {
			continue
		}

		bssid := matches[1]
		signalstr := matches[3]
		signal, err := strconv.Atoi(signalstr)
		if err != nil {
			continue
		}

		aps = append(aps, &AP{
			BSSID:  bssid,
			Signal: signal,
		})
	}

	return aps
}
