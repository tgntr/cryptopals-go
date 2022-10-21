package cryptopals

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_01_04_Detect_Single_Character_Xor(t *testing.T) {
	data, err := os.ReadFile("./01_04_data.txt")
	require.NoError(t, err)

	inputs := strings.Split(string(data), "\n")
	want := "Now that the party is jumping\n"

	var bestScore float64
	var bestResult string
	for _, input := range inputs {
		inputBytes, err := hex.DecodeString(input)
		require.NoError(t, err)

		result, score, _ := decryptCaesarCipher(inputBytes)
		if score > bestScore {
			bestScore = score
			bestResult = result
		}
	}

	have := bestResult
	require.Equal(t, want, have)
}
