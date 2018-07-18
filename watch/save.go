package watch

import (
	firego "gopkg.in/zabawaba99/firego.v1"
	"time"
)

func Save(url string, id string, room Room, timestamp time.Time) error {
	f := firego.New(url, nil)
	f.Child(id).Child("room").Set(room)
	f.Child(id).Child("timestamp").Set(timestamp)
	return nil
}
