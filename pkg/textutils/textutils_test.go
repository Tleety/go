package textutils

import "testing"

func TestFirstN_SendNWithinStringLength_ReturnStringOfNLength(t *testing.T) {
	s := "Hello, world!"
	n := 5
	want := "Hello"
	got := FirstN(s, n)
	if got != want {
		t.Errorf("FirstN(%q, %d) = %q; want %q", s, n, got, want)
	}
}

func TestFirstN_SendNegativeNumber_ReturnEmptyString(t *testing.T) {
	s := "Hello, world!"
	n := -3
	want := ""
	got := FirstN(s, n)
	if got != want {
		t.Errorf("FirstN(%q, %d) = %q; want %q", s, n, got, want)
	}
}

func TestFirstN_SendEmptyString_ReturnEmptyString(t *testing.T) {
	s := ""
	n := 0
	want := ""
	got := FirstN(s, n)
	if got != want {
		t.Errorf("FirstN(%q, %d) = %q; want %q", s, n, got, want)
	}
}

func TestFirstN_SendHigherNumberThanStringLength_ReturnFullString(t *testing.T) {
	s := "Hello, world!"
	n := 50
	want := "Hello, world!"
	got := FirstN(s, n)
	if got != want {
		t.Errorf("FirstN(%q, %d) = %q; want %q", s, n, got, want)
	}
}
