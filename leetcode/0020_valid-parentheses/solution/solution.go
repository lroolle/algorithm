package solution

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if len(stack) > 0 && match(stack[len(stack)-1], c) {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, c)
	}

	return len(stack) == 0
}

func match(c1, c2 rune) bool {
	cs1, cs2 := string(c1), string(c2)
	switch {
	case cs1 == "(" && cs2 == ")":
		return true
	case cs1 == "[" && cs2 == "]":
		return true
	case cs1 == "{" && cs2 == "}":
		return true
	}
	return false
}
