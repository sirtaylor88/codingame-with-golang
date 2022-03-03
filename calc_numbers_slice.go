package main

func Calc(arr []int, n1, n2 int) int {
	sum := 0
	for i, n := range arr {
		if i >= n1 && i <= n2 {
			sum += n
		}
	}
	return sum
}
