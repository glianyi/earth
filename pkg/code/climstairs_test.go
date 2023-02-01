package code

import (
	"fmt"
	"testing"
)

func TestClimStairs(t *testing.T) {
	fmt.Println("test...")
	n := climbStairs(1)
	n1 := climbStairs_v1(1)
	println(n, n1)
	n = climbStairs(2)
	n1 = climbStairs_v1(2)
	println(n, n1)
	n = climbStairs(3)
	n1 = climbStairs_v1(3)
	println(n, n1)
	n = climbStairs(4)
	n1 = climbStairs_v1(4)
	println(n, n1)
	n = climbStairs(5)
	n1 = climbStairs_v1(5)
	println(n, n1)
}
