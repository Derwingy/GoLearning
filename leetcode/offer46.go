package leetcode

func Pow(base int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= base
	}
	return res
}
