package cryptopals

import "strings"

// https://en.wikipedia.org/wiki/Caesar_cipher
func decryptCaesarCipher(cipher []byte) (result string, score float64, key byte) {
	var bestScore float64
	var bestResult string
	var bestKey byte

	for i := 0; i < 256; i++ {
		xor := make([]byte, len(cipher))
		char := byte(i)
		for y := range xor {
			xor[y] = cipher[y] ^ char
		}

		var score float64
		for _, c := range xor {
			charToLower := strings.ToLower(string(c))[0]
			freq, ok := frequencies[charToLower]
			if !ok {
				score = 0
				continue
			}

			score += freq
		}

		score /= float64(len(xor))
		if score > bestScore {
			bestScore = score
			bestResult = string(xor)
			bestKey = byte(char)
		}
	}

	return bestResult, bestScore, bestKey
}

func decryptCaesarCipherRepeatingKey(cipher []byte, key []byte) []byte {
	xor := make([]byte, len(cipher))
	for i := 0; i < len(cipher); i++ {
		key := key[i%len(key)]
		xor[i] = cipher[i] ^ key
	}

	return xor
}

// https://en.wikipedia.org/wiki/Letter_frequency
var frequencies = map[byte]float64{
	'a':  8.2,
	'b':  1.5,
	'c':  2.8,
	'd':  4.3,
	'e':  13,
	'f':  2.2,
	'g':  2,
	'h':  6.1,
	'i':  7,
	'j':  0.15,
	'k':  0.77,
	'l':  4,
	'm':  2.4,
	'n':  6.7,
	'o':  7.5,
	'p':  1.9,
	'q':  0.095,
	'r':  6,
	's':  6.3,
	't':  9.1,
	'u':  2.8,
	'v':  0.98,
	'w':  2.4,
	'x':  0.15,
	'y':  2,
	'z':  0.074,
	' ':  0.1,
	',':  0.1,
	'.':  0.1,
	':':  0.1,
	'\'': 0.1,
	'\n': 0.1,
	'0':  0.1,
	'1':  0.1,
	'2':  0.1,
	'3':  0.1,
	'4':  0.1,
	'5':  0.1,
	'6':  0.1,
	'7':  0.1,
	'8':  0.1,
	'9':  0.1,
}
