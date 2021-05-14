package duh

import "testing"

func TestDuh(t *testing.T) {
	want := "Duh!"
	if got := Duh(); got != want {
		t.Errorf("Duh() = %q, want %q", got, want)
	}
}
