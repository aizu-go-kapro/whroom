package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	xdgbasedir "github.com/zchee/go-xdgbasedir"
	"github.com/zchee/go-xdgbasedir/home"
)

var studentIDRegexp = regexp.MustCompile(`(s|m)\d{7}`)

var configPath = flag.String("config", "", `Location to local config file(default "~/.whroom.toml" or "~/.config/whroom/config.toml")`)

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

func getConfig() (*Config, []string, error) {
	flag.Parse()
	path, ok := getConfigFilePath()
	if !ok {
		return nil, nil, errors.New("could not find config file.") // TODO: add reference
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "could not open %s", path)
	}

	config, err := decodeConfig(file)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse the config file")
	}

	return config, flag.Args(), nil
}

func getConfigFilePath() (string, bool) {
	paths := []string{
		filepath.Join(xdgbasedir.ConfigHome(), "whroom", "config.toml"),
		filepath.Join(home.Dir(), ".whroom.toml"),
		*configPath,
	}
	for _, path := range paths {
		if path == "" {
			continue
		}

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path, true
		}
	}
	return "", false
}
