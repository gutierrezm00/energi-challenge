package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"unicode/utf8"
)

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func isPalindrome(number int) bool {
	// your code goes here
	reversed := Reverse(strconv.Itoa(number))
	chk := reversed == strconv.Itoa(number)

	return chk
}
func main() {
	var number int = rand.Intn(10000)
	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome number", number)
	} else {
		fmt.Printf("%d is NOT a palindrome number", number)
	}
}
