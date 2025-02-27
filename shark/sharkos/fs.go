package sharkos

import (
	"encoding/base64"
	"encoding/hex"
	"io/fs"
	"log"
	"os"
)

type oSFS struct{}

// Open 返回初始文件和用户参数文件的打开操作
func (oSFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func Base64ToBytes(base64Text string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		log.Fatal("error: ", err)
		return []byte{}, err
	}

	return decodedBytes, nil
}

func base64ToHex(base64Text string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		log.Fatal("error: ", err)
		return []byte{}, err
	}

	// 变为hex 16进制
	var result = make([]byte, hex.EncodedLen(len(decodedBytes)))
	hex.Encode(result, decodedBytes)

	return result, nil
}
