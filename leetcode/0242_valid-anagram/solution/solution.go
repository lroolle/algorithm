package solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, r1 := range []rune(s) {
		m[r1] += 1
	}
	for _, r2 := range []rune(t) {
		m[r2] -= 1
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
