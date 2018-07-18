package watch

import (
	"testing"
)

func TestListAP(t *testing.T) {
	airportout := `
     SSID BSSID             RSSI CHANNEL HT CC SECURITY (auth/unicast/group)
   Segers 58:6d:8f:a0:5f:00 -86  36,+1   Y  -- WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP)
    WiFi1 00:18:02:82:b3:00 -84  6       N  BE WEP
    WiFi2 00:1c:df:ed:a9:00 -76  6,-1    Y  US WEP
    WiFi3 06:18:02:82:b3:00 -87  6       N  BE NONE
    WiFi4 5c:35:3b:01:a3:00 -84  1       Y  BE NONE
    WiFi5 5c:35:3b:01:a3:00 -83  1       Y  BE WPA(PSK/TKIP,AES/TKIP) WPA2(PSK/TKIP,AES/TKIP)
   Segers 58:6d:8f:a0:5f:00 -81  11      Y  -- WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP)
`

	cmdstub = airportout

	aps := listAP("abc")

	expected := []*AP{
		{
			BSSID:  "06:18:02:82:b3:00",
			Signal: -87,
		},
		{
			BSSID:  "58:6d:8f:a0:5f:00",
			Signal: -86,
		},
		{
			BSSID:  "5c:35:3b:01:a3:00",
			Signal: -84,
		},
		{
			BSSID:  "00:18:02:82:b3:00",
			Signal: -84,
		},
		{
			BSSID:  "5c:35:3b:01:a3:00",
			Signal: -83,
		},
		{
			BSSID:  "58:6d:8f:a0:5f:00",
			Signal: -81,
		},
		{
			BSSID:  "00:1c:df:ed:a9:00",
			Signal: -76,
		},
	}

	if len(aps) != len(expected) {
		t.Errorf("expected len(aps) is %v but %v", len(expected), len(aps))
	}

	for _, ap := range expected {
		if !containsAP(aps, ap) {
			t.Errorf("expected the %#+v is contained in the return value but don't", *ap)
		}
	}
}

func containsAP(aps []*AP, ap *AP) bool {
	for _, pap := range aps {
		if ap.BSSID == pap.BSSID && ap.Signal == pap.Signal {
			return true
		}
	}
	return false
}
