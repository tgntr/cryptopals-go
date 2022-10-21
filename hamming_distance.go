package cryptopals

import "fmt"

// https://en.wikipedia.org/wiki/Hamming_distance
func calculateHammingDistance(a string, b string) int {
	var aBits string
	// prints string as 8-bit binary
	aBits = fmt.Sprintf("%s%8.b", aBits, []byte(a))

	var bBits string
	bBits = fmt.Sprintf("%s%8.b", bBits, []byte(b))

	distance := 0
	for i := 0; i < len(aBits); i++ {
		if aBits[i] != bBits[i] {
			distance++
		}
	}

	return distance
}
