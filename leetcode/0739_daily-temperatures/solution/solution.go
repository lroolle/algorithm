package solution

func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	// Store [t, i]
	stack := [][]int{}
	var lastT, lastI int
	for i, t := range T {
		if len(stack) > 0 {
			stackhead := stack[len(stack)-1]
			lastT = stackhead[0]

			for lastT < t {
				stack = stack[:len(stack)-1]
				lastI = stackhead[1]
				ret[lastI] = i - lastI

				if len(stack) == 0 {
					break
				}

				stackhead = stack[len(stack)-1]
				lastT = stackhead[0]
			}
		}

		stack = append(stack, []int{t, i})
	}
	return ret
}

func dailyTemperaturesStackOptimize(T []int) []int {
	ret := make([]int, len(T))
	// Store the t index
	stack := []int{}
	for i, t := range T {
		// the "for judgement", no need `if len(stack) > 0` !
		// the `lastT` is always `stack[len(stack)-1]`
		for len(stack) > 0 && t > T[stack[len(stack)-1]] {
			lastI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[lastI] = i - lastI
		}
		stack = append(stack, i)
	}
	return ret
}
