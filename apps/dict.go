package apps

import (
	"os"
	"path/filepath"
)

func (a *App) GetDictContent() (string, error) {
	workHome, _ := os.Getwd()
	DictFilePath := filepath.Join(workHome, "static/dict/"+"cain.txt")
	//读取文件内容，按行分割
	content, err := os.ReadFile(DictFilePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
