package scrapper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

func generateKey(passPhrase, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	return pbkdf2.Key(passPhrase, salt, iter, keyLen, sha1.New)
}
func buildAesPwd(keyPass, plain string) {

	key, err := hex.DecodeString(keyPass)
	if err != nil {
		log.Fatal(err)
	}
	plaintext := []byte(plain)

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}

type AesUtil struct {
	KeySize        int
	IterationCount int
}

func (au AesUtil) generateKey(salt, passPhrase) (key []byte) {
	// func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte
	return pbkdf2.Key(passPhrase, salt, au.IterationCount, au.KeySize, sha1.New)
}
