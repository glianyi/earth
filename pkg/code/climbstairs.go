package code

func climbStairs(n int) int {
	if n == 2 || n == 1 || n == 0 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs_v1(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	pre, next := 1, 2
	tmp := 0
	for ; n > 2; n-- {
		tmp = next
		next = pre + next
		pre = tmp
	}
	return next
}
