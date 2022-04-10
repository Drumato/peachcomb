package strparse

// hasPrefix check s starts with t
func hasPrefix(s []rune, t []rune) bool {
	if len(s) < len(t) {
		return false
	}

	for i := range t {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}
