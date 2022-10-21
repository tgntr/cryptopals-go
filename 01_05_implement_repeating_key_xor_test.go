package cryptopals

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_01_05_Implement_Repeating_Key_Xor(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" + "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	decrypted := decryptCaesarCipherRepeatingKey([]byte(input), []byte(key))
	have := hex.EncodeToString(decrypted)
	require.Equal(t, want, have)
}
