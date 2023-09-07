package utils

import (
	"io/ioutil"
	"os"
)

// ReadFile 读取文件内容
func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile 写入文件内容
func WriteFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)
	return err
}
