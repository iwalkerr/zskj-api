package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	// 加密
	dd, err := AESEncrypt([]byte("1kkkkkk"), []byte("123456781234567812345678"))
	if err != nil {
		fmt.Println("1111", err)
		return
	}

	fmt.Println(string(dd))

	dd, err = AESDecrypt(dd, []byte("123456781234567812345678"))
	if err != nil {
		fmt.Println("222", err)
		return
	}
	fmt.Println(string(dd))
}
