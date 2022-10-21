package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_01_01_Hex_To_Base64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	inputBytes, err := hex.DecodeString(input)
	require.NoError(t, err)

	have := base64.StdEncoding.EncodeToString(inputBytes)
	require.Equal(t, want, have)
}

func Test_01_02_Fixed_Xor(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"

	input1Bytes, err := hex.DecodeString(input1)
	require.NoError(t, err)

	input2Bytes, err := hex.DecodeString(input2)
	require.NoError(t, err)

	xor := make([]byte, len(input1Bytes))
	for i := range xor {
		xor[i] = input1Bytes[i] ^ input2Bytes[i]
	}

	have := hex.EncodeToString(xor)
	require.Equal(t, want, have)
}

func Test_01_03_Single_Byte_Xor_Cipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"

	inputBytes, err := hex.DecodeString(input)
	require.NoError(t, err)

	have, _, _ := decryptCaesarCipher(inputBytes)
	require.Equal(t, want, have)
}

func Test_01_04_Detect_Single_Character_Xor(t *testing.T) {
	data, err := os.ReadFile("data/01_04.txt")
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

func Test_01_05_Implement_Repeating_Key_Xor(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" + "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	decrypted := decryptCaesarCipherRepeatingKey([]byte(input), []byte(key))
	have := hex.EncodeToString(decrypted)
	require.Equal(t, want, have)
}

func Test_01_06_Break_Repeating_Key_Xor(t *testing.T) {
	data, err := os.ReadFile("data/01_06.txt")
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
