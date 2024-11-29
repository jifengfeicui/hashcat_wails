package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func CalculateMD5(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建一个新的 MD5 哈希对象
	hash := md5.New()

	// 将文件内容拷贝到 MD5 哈希对象中
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// 获取文件的 MD5 哈希值
	md5Sum := hash.Sum(nil)

	// 返回 MD5 值的十六进制字符串
	return fmt.Sprintf("%x", md5Sum), nil
}
