package main

func GetNextSequenceNumber(n int) int {
	nums := []int{}

	nums = append(nums, n)

	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
	}

	sum := 0

	for _, digit := range nums {
		sum = sum + digit
	}
	return sum
}

func FindJointPointOfSequences(seq1, seq2 int) int {
	for seq1 != seq2 {
		if seq1 < seq2 {
			seq1 = GetNextSequenceNumber(seq1)
		} else if seq2 < seq1 {
			seq2 = GetNextSequenceNumber(seq2)
		}
	}
	return seq1
}
