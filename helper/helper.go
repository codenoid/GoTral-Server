/*
* github.com/codenoid - Developer
* code source : - https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
*               - https://golang.org/pkg/crypto/cipher
*               - https://github.com/codenoid/GoTral
*
 */
package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// createHash : create md5 hash and return as string
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt : encrypt given data with passphrase
// Load your secret key from a safe place and reuse it across multiple
// Seal/Open calls, this library actually doesn't need
func Encrypt(data []byte, passphrase string) ([]byte, error) {
	// create aes.NewCipher from hashed md5 passphrase
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	//  NewGCM returns the given 128-bit, block cipher wrapped in
	// Galois Counter Mode with the standard nonce length.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// initialize slice with length of nonce that must be passed to Seal and Open.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}
