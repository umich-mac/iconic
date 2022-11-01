package main

import "testing"

func TestHashing(t *testing.T) {
	want := "2f5bbd5c5147380eecd94daa0cf0681c30ca43803f10868a2b9a2aedd26b0f21"
	if got := HashFile("test.png"); got != want {
		t.Errorf("HashFile('test.png') = %q, want %q", got, want)
	}
}
