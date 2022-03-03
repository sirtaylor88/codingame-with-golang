package main

import (
	"fmt"
	"math"
)

func ComputeClosestToZero(ts []int) int {
	// Write your code here
	// fmt.Fprintln(os.Stderr, "Debug messages...")

	if len(ts) == 0 {
		return 0
	}

	closest := 0
	for _, t := range ts {
		if closest == 0 {
			closest = t
		} else if t > 0 && t <= int(math.Abs(float64(closest))) {
			closest = t
		} else if t < 0 && -t < int(math.Abs(float64(closest))) {
			closest = t
		}
	}
	return closest

}

/* Ignore and do not change the code below */
func ClosestToZero() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Buffer(make([]byte, 1000000), 1000000)
	// var inputs []string

	// var n int
	// scanner.Scan()
	// fmt.Sscan(scanner.Text(), &n)

	// scanner.Scan()
	// inputs = strings.Split(scanner.Text(), " ")
	// ts := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	ts[i], _ = strconv.Atoi(inputs[i])
	// }
	// stdoutWriter := os.Stdout
	// os.Stdout = os.Stderr
	// solution := ComputeClosestToZero(ts)
	// os.Stdout = stdoutWriter
	// fmt.Println(solution)

	ts := []int{5, -5}
	fmt.Println("Closest value of [5 -5] to zero is", ComputeClosestToZero(ts))

	ts = []int{10}
	fmt.Println("Closest value of [10] to zero is", ComputeClosestToZero(ts))

	ts = []int{-10, -10}
	fmt.Println("Closest value of [-10 -10] to zero is", ComputeClosestToZero(ts))

	ts = []int{-273}
	fmt.Println("Closest value of [-273] to zero is", ComputeClosestToZero(ts))

	ts = []int{18, -12, -23, 15}
	fmt.Println("Closest value of [18 -12 -23 15] to zero is", ComputeClosestToZero(ts))

	ts = []int{-100, -10, -75}
	fmt.Println("Closest value of [-100 -10 -75] to zero is", ComputeClosestToZero(ts))

	ts = []int{}
	fmt.Println("Closest value of [] to zero is", ComputeClosestToZero(ts))

	ts = []int{4, 0}
	fmt.Println("Closest value [4 0] to zero is", ComputeClosestToZero(ts))
}
