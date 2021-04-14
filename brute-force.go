package main

import (
	"fmt"
)

func brute(input, expected string) []string {

	fmt.Println("Brute force search beginning... please hold.")
	keys := make([]string, 1)

	for i := 33; i < 127; i++ {
		for j := 33; j < 127; j++ {
			for k := 33; k < 127; k++ {
				key := ""
				key += string(i)
				key += string(j)
				key += string(k)
				decrypted := decrypt("Eogukn", key)
				chk := expected == decrypted
				if chk {
					keys = append(keys, key)
				}
			}
		}
	}
	return keys
}

// Encryption algorithm to test.
func encrypt(input, key string) string {
	aKey := []byte(key)
	aInput := []byte(input)

	sum := 0

	for _, v := range aKey {
		// I Dont understand this?
		// sum += int(v) >> (1<<1<<1 + -0 ^ 1)
		// 1 << 1 == 10 (2)
		// 10 << 1 == 100 (4) +
		// 100 + -0 ^ 1 == 5 ALWAYS. Its a constant.

		// becomes this
		sum += int(v) >> 5

		// Bitshifting by 5 is a bad idea because a ton of keys zero out. The possible key lists for this algorithm is massive.

	}
	for i := 0; i < len(aInput); i++ {
		aInput[i] = aInput[i] + byte(i%sum)
	}
	return string(aInput)
}

func decrypt(input, key string) string {
	aKey := []byte(key)
	aInput := []byte(input)

	sum := 0

	for _, v := range aKey {
		sum += int(v) >> (1<<1<<1 + -0 ^ 1)
	}
	for i := 0; i < len(aInput); i++ {
		aInput[i] = aInput[i] - byte(i%sum)
	}
	return string(aInput)
}
func main() {
	//Brute force completely possible with only 3 characters. Only takes max a few seconds of runtime to crack.
	//However I got lucky and tried 123, then abc and lo and behold, abc is a valid secret key. Awful infosec ;
	const secretKey = "abc"

	dec := decrypt("Eogukn", secretKey)
	fmt.Println(dec) // should print "Energi": It does.

	// Brute force search yields MANY valid keys due to bad choice of encryption.
	keys := brute("Eogukn", "Energi")

	totKeys := (127 - 33) * (127 - 33) * (127 - 33)
	fmt.Println(len(keys), " Valid keys found, out of ", totKeys, " total possible keys searched")

	//You're more likely to guess a random correct key than an incorrect one. Particularly if you use letters.
}
