package base

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// AES CBC 加密
func AesEncryptCBC(origData, key []byte) (crypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			crypted = nil
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted = make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return
}

// AES CBC 解密
func AesDecryptCBC(crypted, key []byte) (origData []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			origData = nil
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return
}

// AES CBC 加密后做一次Base64加密
func AesEncryptCBCWithBase64(origData, key []byte) (string, error) {
	cbc, err := AesEncryptCBC(origData, key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cbc), nil
}

// 先Base64解密，再AES CBC解密
func AesDecryptCBCWithBase64(cryptedWithBase64 string, key []byte) ([]byte, error) {
	bytesCbc, err := base64.StdEncoding.DecodeString(cryptedWithBase64)
	if err != nil {
		return nil, err
	}

	return AesDecryptCBC(bytesCbc, key)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == 0 {
		return ciphertext
	}

	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	for i := len(origData) - 1; i >= 0; i-- {
		if origData[i] != 0 {
			return origData[:i+1]
		}
	}

	return nil
}
