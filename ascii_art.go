package main

// printChar allows to display one and only one ASCII character from A to Z (inclusive) on several lines and columns
// (graphic representation called "ASCII Art").

// We want to perform the operation in the other direction: obtain a character from its representation graphic.
// Implement the scanChar (str) method so that it returns the character c associated with its representation graph
// (i.e. str = printChar (c)). If str does not match a character between A and Z (inclusive),
// then scanChar should return the character '?'

// func scanChar(str string) string {
// 	for char := 'A'; char <= 'Z'; char++ {
// 		if printChar(string(char)) == str {
// 			return string(char)
// 		}
// 	}
// 	return "?"
// }
