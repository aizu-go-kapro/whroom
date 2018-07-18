package watch

import (
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/BurntSushi/toml"
)

var studentIDRegexp = regexp.MustCompile(`(s|m)\d{7}`)

type Config struct {
	FirebaseURL   string
	StudentID     string
	WifiInterface string
	Duration      time.Duration
}

type ConfigUnmarshaler struct {
	FirebaseURL   string   `toml:"firebase_url"`
	StudentID     string   `toml:"student_id"`
	WifiInterface string   `toml:"wifi_interface"`
	Duration      duration `toml:"duration"`
}

type duration time.Duration

func (d *duration) UnmarshalText(data []byte) error {
	dur, err := time.ParseDuration(string(data))
	if err != nil {
		return err
	}
	*d = duration(dur)
	return nil
}

func decodeConfig(r io.Reader) (*Config, error) {
	// Set the default values.
	cu := &ConfigUnmarshaler{
		Duration: duration(1 * time.Minute),
	}

	if _, err := toml.DecodeReader(r, &cu); err != nil {
		return nil, err
	}

	if cu.FirebaseURL == "" {
		return nil, errors.New("the firebase_url field is not set")
	}
	if !studentIDRegexp.MatchString(cu.StudentID) {
		return nil, errors.New("the student_id field has invalid value or empty")
	}

	return &Config{
		FirebaseURL:   cu.FirebaseURL,
		StudentID:     cu.StudentID,
		WifiInterface: cu.WifiInterface,
		Duration:      time.Duration(cu.Duration),
	}, nil
}
