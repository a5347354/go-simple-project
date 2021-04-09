package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func AES256Decrypter(ciphertext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	block, err := aes.NewCipher(bKey)

	mode := cipher.NewCBCDecrypter(block, bIV)
	source, err := hex.DecodeString(ciphertext)
	if err != nil {
		panic(err)
	}
	var dst []byte = make([]byte, len(source))
	mode.CryptBlocks(dst, source)
	origin := string(PKCS5Unpadding(dst))
	return origin
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AES256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize)
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, pad...)
}
