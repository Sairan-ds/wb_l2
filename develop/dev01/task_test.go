package main

import (
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	got := CurrentTime()
	want := time.Since(got)
	if want > 50*time.Millisecond {
		t.Errorf("got %q, wanted %q", got, want)
	}
}