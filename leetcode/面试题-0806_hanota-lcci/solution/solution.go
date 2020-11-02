package solution

func hanota(A []int, B []int, C []int) []int {
	var move func(int, *[]int, *[]int, *[]int)
	move = func(n int, from, mid, to *[]int) {
		if n == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}

		move(n-1, from, to, mid)
		move(1, from, mid, to)
		move(n-1, mid, from, to)
	}

	move(len(A), &A, &B, &C)
	return C
}
