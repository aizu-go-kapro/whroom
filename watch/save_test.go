package watch

import (
	"testing"
	"time"

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
	s, ok := v.(string)
	if !ok {
		t.Fatalf("mismatch type of value: %#+v", v)
	}

	if s != "M1" {
		t.Fatalf("expected M1 but %s", s)
	}
}
