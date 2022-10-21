package cryptopals

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_01_03_Single_Byte_Xor_Cipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"

	inputBytes, err := hex.DecodeString(input)
	require.NoError(t, err)

	have, _, _ := decryptCaesarCipher(inputBytes)
	require.Equal(t, want, have)
}
