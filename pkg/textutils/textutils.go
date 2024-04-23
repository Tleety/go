package textutils

func FirstN(s string, n int) string {
	if n < 0 {
		return ""
	} else if n > len(s) {
		return s
	} else {
		return s[:n]
	}
}

func LastN(s string, n int) string {
	if n < 0 {
		return ""
	} else if n > len(s) {
		return s
	} else {
		return s[len(s)-n:]
	}
}
