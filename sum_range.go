package main

func SumRange(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		if n >= 10 && n <= 100 {
			sum += n
		}
	}
	return sum
}
