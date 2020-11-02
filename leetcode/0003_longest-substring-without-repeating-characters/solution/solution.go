package solution

// Hash Table to hold the index of
func lengthOfLongestSubstring(s string) int {
	var ret, start int
	var table = make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if index, ok := table[s[i]]; ok && index >= start {
			start = index + 1
		}
		sublength := i - start + 1
		if sublength > ret {
			ret = sublength
		}
		table[s[i]] = i
	}

	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLongestSubstring1(s string) int {
	var start, ret int
	table := make(map[rune]int)

	for i, c := range s {
		if index, ok := table[c]; ok {
			ret = max(ret, i-start)
			start = max(start, index+1)
		}
		table[c] = i
	}
	ret = max(ret, len(s)-start)
	return ret
}

func lengthOfLongestSubstring2(s string) int {
	var freq [256]byte
	size := len(s)
	result, left, right := 0, 0, 0

	for left < size {
		if right < size && freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}

		sublength := right - left
		if sublength > result {
			result = sublength
		}
	}
	return result
}
