package libs

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	configs "catdogs.club/back-end/configs/common"
)

func PKCS7Padding(org []byte, blockSize int) []byte {
	pad := blockSize - len(org)%blockSize
	padArr := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(org, padArr...)
}

func PKCS7UnPadding(org []byte) []byte {
	l := len(org)
	pad := org[l-1]
	return org[:l-int(pad)]
}

func AESDecrypt(str string) string {
	block, _ := aes.NewCipher([]byte(configs.AesKey))
	blockMode := cipher.NewCBCDecrypter(block, []byte(configs.AesKey))
	old, _ := hex.DecodeString(str)
	org := make([]byte, len(old))
	blockMode.CryptBlocks(org, []byte(old))
	org = PKCS7UnPadding(org)
	return string(org)
}

func AESEncrypt(str string) string {
	block, _ := aes.NewCipher([]byte(configs.AesKey))
	org := PKCS7Padding([]byte(str), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(configs.AesKey))
	cryted := make([]byte, len(org))
	blockMode.CryptBlocks(cryted, org)
	return hex.EncodeToString(cryted)
}
