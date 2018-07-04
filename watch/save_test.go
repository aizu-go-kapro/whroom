package watch

import (
	"testing"
	"time"

	"fmt"
	"github.com/zabawaba99/firego/firetest"
)

func TestSave(t *testing.T) {
	server := firetest.New()
	defer server.Close()

	server.Start()

	tm := time.Date(2018, 1, 1, 1, 0, 0, 0, time.UTC)
	if err := Save(server.URL, "s1220004", M1, tm); err != nil {
		t.Fatal(err)
	}
	v := server.Get("s1220004")
	s, ok := v.(map[string]interface{})
	if !ok {
		t.Fatalf("mismatch type of value: %#+v", v)
	}

	if s["room"] != "M1" {
		t.Fatalf("expected M1 but %s", s)
	}
	if s["timestamp"] != "2018-01-01T01:00:00Z" {
		t.Fatalf("Not True")
	}
}
