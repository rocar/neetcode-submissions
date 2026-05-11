func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}

		for r > l && !isAlphaNum(s[r]) {
			r--
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return ('A' <= c && c <='Z') ||
	    ('a' <= c && c <= 'z') ||
		('0' <= c && c <= '9')
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 32
	}
	return c
}