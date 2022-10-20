package cryptopals

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Single-byte XOR cipher

// The hex encoded string:
// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

// ... has been XOR'd against a single character. Find the key, decrypt the message.

// You can do this by hand. But don't: write code to do it for you.

// How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
func Test_01_03_Single_Byte_Xor_Cipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"

	inputBytes, err := hex.DecodeString(input)
	require.NoError(t, err)

	var bestScore float64
	var bestResult []byte

	for i := 0; i < 256; i++ {
		xor := make([]byte, len(inputBytes))
		char := byte(i)
		for i := range xor {
			xor[i] = inputBytes[i] ^ char
		}

		var score float64
		for _, c := range xor {
			charToLower := strings.ToLower(string(c))[0]
			freq, ok := frequencies[charToLower]
			if !ok {
				score = 0
				break
			}

			score += freq
		}

		if score > bestScore {
			bestScore = score
			bestResult = xor
		}
	}

	have := string(bestResult)
	require.Equal(t, want, have)
}

var frequencies = map[byte]float64{
	'a':  8.2,
	'b':  1.5,
	'c':  1.5,
	'd':  1.5,
	'e':  1.5,
	'f':  1.5,
	'g':  1.5,
	'h':  1.5,
	'i':  1.5,
	'j':  1.5,
	'k':  1.5,
	'l':  1.5,
	'm':  1.5,
	'n':  1.5,
	'o':  1.5,
	'p':  1.5,
	'q':  1.5,
	'r':  1.5,
	's':  1.5,
	't':  1.5,
	'u':  1.5,
	'v':  1.5,
	'w':  1.5,
	'x':  1.5,
	'y':  1.5,
	'z':  1.5,
	' ':  0.1,
	',':  0.1,
	'.':  0.1,
	'\'': 0.1,
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
