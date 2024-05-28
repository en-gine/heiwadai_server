package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"server/infrastructure/env"
)

var EncryptKey = env.GetEnv(env.EncryptKey)

// Encrypt encrypts a plain text string using AES encryption.
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(EncryptKey))
	if err != nil {
		return "", err
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	paddedPlainText := pad([]byte(plainText), aes.BlockSize)
	cipherText := make([]byte, len(paddedPlainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedPlainText)

	// Encode IV and ciphertext to base64
	encodedIV := base64.StdEncoding.EncodeToString(iv)
	encodedCipherText := base64.StdEncoding.EncodeToString(cipherText)

	return fmt.Sprintf("%s:%s", encodedIV, encodedCipherText), nil
}

// Decrypt decrypts a base64 encoded ciphertext string using AES encryption.
func Decrypt(encryptedText string) (string, error) {
	// Split the IV and the encrypted text
	parts := split(encryptedText, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid encrypted text format")
	}

	// Decode IV and ciphertext from base64
	iv, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}
	cipherText, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(EncryptKey))
	if err != nil {
		return "", err
	}

	// Decrypt the ciphertext
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)
	unpaddedPlainText, err := unpad(plainText, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(unpaddedPlainText), nil
}

// Split a string by a delimiter and return the parts
func split(s, delimiter string) []string {
	return append([]string{}, strings.Split(s, delimiter)...)
}

// pad adds padding to the plaintext to make it a multiple of the block size
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// unpad removes padding from the plaintext
func unpad(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, fmt.Errorf("plain text is empty")
	}
	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, fmt.Errorf("plain text has invalid padding")
	}
	return src[:(length - unpadding)], nil
}
