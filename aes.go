package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
)

// Sha256Key sha256 加密
func Sha256Key(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	newKey := h.Sum(nil)
	return newKey
}

// PKCS7Padding 填充数据
func PKCS7Padding(ciphertext []byte) []byte {
	bs := aes.BlockSize
	padding := bs - len(ciphertext)%bs
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

// PKCS7UnPadding 放出数据
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密
func AesEncrypt(origData, key string) (string, error) {
	newKey := Sha256Key(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	newOrigData := []byte(origData)
	newOrigData = PKCS7Padding(newOrigData)
	blockMode := cipher.NewCBCEncrypter(block, newKey[:16])
	crypted := make([]byte, len(newOrigData))
	blockMode.CryptBlocks(crypted, newOrigData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt 解密
func AesDecrypt(crypted, key string) (string, error) {
	newKey := Sha256Key(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	newCrypted, _ := base64.StdEncoding.DecodeString(crypted)
	blockMode := cipher.NewCBCDecrypter(block, newKey[:16])
	origData := make([]byte, len(newCrypted))
	blockMode.CryptBlocks(origData, newCrypted)
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}
