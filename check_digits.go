package main

import (
	"strconv"
	"strings"
)

func ComputeCheckDigits(identificationNumber string) int {
	digits := strings.Split(identificationNumber, "")
	sum_even_index := 0
	sum_odd_index := 0
	for i, digit := range digits {
		n, _ := strconv.Atoi(digit)
		if i%2 == 0 {
			sum_even_index += n
		} else {
			sum_odd_index += n
		}
	}
	checksum := 3*sum_even_index + sum_odd_index
	last_digit := checksum % 10
	if last_digit == 0 {
		return 0
	} else {
		return 10 - last_digit
	}
}
