package main

import "testing"

func TestGetRoom(t *testing.T) {
	room, ok := getRoom("6c:dd:30:37:39:ab")
	if !ok {
		t.Fatal("expected any room can be got but didn't")
	}
	if room != M7 {
		t.Errorf("expected M7 but %v", room)
	}
}
