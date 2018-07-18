package main

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeConfig(t *testing.T) {
	tests := map[string]struct {
		confstr string
		config  *Config
		wantErr bool
	}{
		"1": {
			confstr: `
firebase_url = "https://test.firebaseio.com"
student_id = "s1230004"
wifi_interface = "wlp2s0"
duration = "1s"
`,
			config: &Config{
				FirebaseURL:   "https://test.firebaseio.com",
				StudentID:     "s1230004",
				WifiInterface: "wlp2s0",
				Duration:      1 * time.Second,
			},
			wantErr: false,
		},
		"2": {
			confstr: `
firebase_url = "https://test.firebaseio.com"
student_id = "s1230004"
wifi_interface = "wlp2s0"
`,
			config: &Config{
				FirebaseURL:   "https://test.firebaseio.com",
				StudentID:     "s1230004",
				WifiInterface: "wlp2s0",
				Duration:      1 * time.Minute,
			},
			wantErr: false,
		},
		"3": {
			confstr: `
firebase_url = "https://test.firebaseio.com"

student_id = "f1230004"
wifi_interface = "wlp2s0"
`,
			config: &Config{
				FirebaseURL:   "https://test.firebaseio.com",
				StudentID:     "s1230004",
				WifiInterface: "wlp2s0",
				Duration:      1 * time.Minute,
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			conf, err := decodeConfig(strings.NewReader(test.confstr))
			if (err != nil) != test.wantErr {
				t.Fatalf("wantErr is %v but err: %v", test.wantErr, err)
			}
			if !test.wantErr {
				if diff := cmp.Diff(conf, test.config); diff != "" {
					t.Fatalf("Config has diff: %s", diff)
				}
			}
		})
	}
}
