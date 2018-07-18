package watch

import (
	"strconv"
	"strings"
	"time"
)

const minimumThreshold = -100

func listAP(ifname string) []*AP {
	c := make(chan map[string]map[string]int)
	go iw(c, ifname)

	var aps []*AP
	for v := range c {
		m, ok := v["wifi"]
		if !ok {
			continue
		}

		for bssid, signal := range m {
			aps = append(aps, &AP{
				BSSID:  bssid,
				Signal: signal,
			})
		}
	}

	return aps
}

func iw(out chan map[string]map[string]int, ifname string) {
	s, _ := RunCommand(10*time.Second, "/sbin/iw dev "+ifname+" scan -u")
	name := ""
	signal := 0
	datas := make(map[string]map[string]int)
	datas["wifi"] = make(map[string]int)
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
	close(out)
}
