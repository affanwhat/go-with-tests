package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Affan")
	want := "Hello, Affan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}