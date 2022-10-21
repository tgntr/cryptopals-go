package cryptopals

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

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
