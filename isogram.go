package isogram

import "unicode"

// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

// IsIsogram
func IsIsogram(in string) bool {

	if len(in) == 0 {
		return true
	}

	// in = string.ToUpper(in)

	chars := make(map[rune]bool)
	for _, ch := range in {
		ch = unicode.ToUpper(ch)
		if !chars[ch] {
			chars[ch] = true
		} else {
			return false
		}
	}

	return true
}

/*
FIXED
- isogram_test.go:8: FAIL: word with duplicated character in mixed case
- Word "Alphabet", expected false, got true

TODO
- isogram_test.go:8: FAIL: isogram with duplicated hyphen
- Word "six-year-old", expected true, got false
*/
