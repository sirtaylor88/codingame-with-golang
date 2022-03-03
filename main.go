package main

import "fmt"

func main() {
	fmt.Println("Largest number is", FindLargest([]int{10, 5, 12, 8})) // 12
	ClosestToZero()
	fmt.Println("Joint Point of Sequences 471 480", FindJointPointOfSequences(471, 480))
	// scanChar("A")
	fmt.Println("Checksum of check digits", ComputeCheckDigits("39847")) // 3
	fmt.Println("Encoded plaintext", Encode("aabaa"))                    //2a1b2a
	fmt.Println("Encoded plaintext", Encode("aaaabcccaaa"))              // 4a1b3c3a
	fmt.Println("SumRange", SumRange([]int{1, 20, 3, 10, -2, 100}))      // 130

	arr := []int{0, 1, 2, 3, 4, 5, 3}
	fmt.Println(Calc(arr, 0, 1)) // 1
	fmt.Println(Calc(arr, 0, 5)) // 15
	fmt.Println(Calc(arr, 0, 0)) // 0
	fmt.Println(Calc(arr, 0, 6)) // 18

	fmt.Println(Sqrt(4))  // 2 <nil>
	fmt.Println(Sqrt(9))  // 3 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2

	fmt.Println(Pic(3, 3)) // [[0 0 1] [0 1 1] [1 1 2]]
	WaitChannels()
	fmt.Println(WordCount("I am learning Go!")) // map[Go!:1 I:1 am:1 learning:1]

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
