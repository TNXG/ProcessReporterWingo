package tools

import (
	"os"
)

// Getwd函数将获取当前的工作目录
func Getwd() (string, error) {
	return os.Getwd()
}
