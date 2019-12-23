package isogram

import "unicode"

// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

// IsIsogram is a string made of unique letters, handle mixed case, ignore spaces and hyphens
func IsIsogram(in string) bool {

	// in = strings.ToUpper(in) // unicode.ToUpper is faster
	chars := make(map[rune]bool)
	for _, ch := range in {

		if ch == '-' || ch == ' ' {
			continue
		}

		ch = unicode.ToUpper(ch)
		if !chars[ch] {
			chars[ch] = true
		} else {
			return false
		}
	}

	return true
}
