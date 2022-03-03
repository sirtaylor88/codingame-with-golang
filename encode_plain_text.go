package main

import (
	"strconv"
	"strings"
)

func Encode(plaintext string) string {
	// aabaa -> 2ab2a
	temp := plaintext[0:1]
	count := 1
	chrs := strings.Split(plaintext, "")
	chrs = chrs[1:]

	var encodedString string
	for _, ch := range chrs {
		if ch == temp {
			count = count + 1

		} else {
			encodedString += strconv.Itoa(count) + temp
			temp = ch
			count = 1
		}
	}
	encodedString += strconv.Itoa(count) + temp
	return encodedString
}
