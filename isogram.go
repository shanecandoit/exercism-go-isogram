package isogram

// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

// IsIsogram
func IsIsogram(in string) bool {

	if len(in) == 0 {
		return true
	}

	chars := make(map[rune]bool)
	for _, ch := range in {

		if !chars[ch] {
			chars[ch] = true
		} else {
			return false
		}
	}

	return true
}
