package textutils

import "testing"

func TestFirstN(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"SendNWithinStringLength", "Hello, world!", 5, "Hello"},
		{"SendNegativeNumber", "Hello, world!", -3, ""},
		{"SendEmptyString", "", 0, ""},
		{"SendHigherNumberThanStringLength", "Hello, world!", 50, "Hello, world!"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FirstN(c.s, c.n)
			if got != c.want {
				t.Errorf("FirstN(%q, %d) = %q; want %q", c.s, c.n, got, c.want)
			}
		})
	}
}

func TestLastN(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"SendNWithinStringLength", "Hello, world!", 5, "world"},
		{"SendNegativeNumber", "Hello, world!", -3, ""},
		{"SendEmptyString", "", 0, ""},
		{"SendHigherNumberThanStringLength", "Hello, world!", 50, "Hello, world!"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := LastN(c.s, c.n)
			if got != c.want {
				t.Errorf("LastN(%q, %d) = %q; want %q", c.s, c.n, got, c.want)
			}
		})
	}
}
