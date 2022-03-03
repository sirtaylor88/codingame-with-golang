package main

import (
	"sort"
)

//"fmt" // Use fmt.Println(...) to debug your code

// FindLargest returns the largest number in the numbers array
func FindLargest(numbers []int) int {
	// Your code goes here
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}
