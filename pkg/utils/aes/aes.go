package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//解密
func AESDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

//去补码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

//加密
func AESEncrypt(origData, key []byte) ([]byte, error) {
	//获取block块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//补码
	origData = PKCS7Padding(origData, block.BlockSize())
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//补码
func PKCS7Padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}
