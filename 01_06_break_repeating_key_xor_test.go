package cryptopals

import (
	"encoding/base64"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_01_06_Break_Repeating_Key_Xor(t *testing.T) {
	data, err := os.ReadFile("./01_06_data.txt")
	require.NoError(t, err)

	input, err := base64.StdEncoding.DecodeString(string(data))
	require.NoError(t, err)

	want := "Terminator X: Bring the noise"

	distances := map[float64]int{}

	// 1. Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
	for keysize := 2; keysize <= 40; keysize++ {
		//3. For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
		a := string(input[keysize : keysize*2])
		b := string(input[keysize*2 : keysize*3])
		distance := float64(calculateHammingDistance(a, b)) / float64(keysize)
		distances[distance] = keysize
	}

	// 4. The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.
	distancesSorted := []float64{}
	for k := range distances {
		distancesSorted = append(distancesSorted, k)
	}
	sort.Float64s(distancesSorted)

	var bestScore float64
	var bestKey string
	for _, d := range distancesSorted[0:3] {
		keysize := distances[d]
		blocks := make([][]byte, keysize)

		// 5. Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
		for i := 0; i < len(input); i += keysize {
			currentBlock := input[i : i+keysize]

			// 6. Now transpose the blocks: make a currentBlock that is the first byte of every currentBlock, and a currentBlock that is the second byte of every currentBlock, and so on.
			for y := range currentBlock {
				blocks[y] = append(blocks[y], currentBlock[y])
			}
		}

		// 7.Solve each block as if it was single-character XOR. You already have code to do this.
		key := make([]byte, len(blocks))
		var score float64
		for i := range blocks {
			_, s, k := decryptCaesarCipher(blocks[i])
			score += s
			key[i] = k
		}

		// 8. For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
		if score > bestScore {
			bestScore = score
			bestKey = string(key)
		}
	}

	have := bestKey
	require.Equal(t, want, have)
}

func Test_Calculate_Hamming_Distance(t *testing.T) {
	// 2. Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between "this is a test" and  "wokka wokka!!!" is 37. Make sure your code agrees before you proceed.
	require.Equal(t, 37, calculateHammingDistance("wokka wokka!!!", "this is a test"))
}
