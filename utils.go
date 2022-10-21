package cryptopals

import (
	"crypto/aes"
	"fmt"
	"strings"
)

// https://en.wikipedia.org/wiki/Caesar_cipher
func decryptSingleByteXOR(msg []byte) (result string, score float64, key byte) {
	var bestScore float64
	var bestResult string
	var bestKey byte

	for i := 0; i < 256; i++ {
		res := make([]byte, len(msg))
		char := byte(i)
		for y := range res {
			res[y] = msg[y] ^ char
		}

		var score float64
		for _, c := range res {
			charToLower := strings.ToLower(string(c))[0]
			freq, ok := frequencies[charToLower]
			if !ok {
				score = 0
				continue
			}

			score += freq
		}

		score /= float64(len(res))
		if score > bestScore {
			bestScore = score
			bestResult = string(res)
			bestKey = byte(char)
		}
	}

	return bestResult, bestScore, bestKey
}

// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
func encryptRepeatingKeyXOR(msg []byte, key []byte) []byte {
	res := make([]byte, len(msg))
	for i := 0; i < len(msg); i++ {
		key := key[i%len(key)]
		res[i] = msg[i] ^ key
	}

	return res
}

// https://en.wikipedia.org/wiki/Hamming_distance
func calculateHammingDistance(a []byte, b []byte) int {
	var aBits string
	// prints string as 8-bit binary
	aBits = fmt.Sprintf("%s%8.b", aBits, a)

	var bBits string
	bBits = fmt.Sprintf("%s%8.b", bBits, b)

	var distance int
	for i := 0; i < len(aBits); i++ {
		if aBits[i] != bBits[i] {
			distance++
		}
	}

	return distance
}

func decryptAesEcb(msg []byte, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	res := make([]byte, len(msg))
	for i := 0; i < len(msg); i += ecbBlockSize {
		end := i + ecbBlockSize
		cipher.Decrypt(res[i:end], msg[i:end])
	}

	return res, nil
}

// msg is encrypted in ECB mode if it has repeating 16 byte chunks
func isEncryptedAesEcb(msg []byte) bool {
	res := [][]byte{}
	for i := 0; i < len(msg); i += ecbBlockSize {
		end := i + ecbBlockSize
		current := msg[i:end]
		for _, r := range res {
			if calculateHammingDistance(r, current) == 0 {
				return true
			}
		}
		res = append(res, current)
	}

	return false
}

const (
	ecbBlockSize = 16
)

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
