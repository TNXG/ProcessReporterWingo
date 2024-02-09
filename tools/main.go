package tools

import (
	"bytes"
	"os"
	"os/exec"
)

// Getwd函数将获取当前的工作目录
func Getwd() (string, error) {
	return os.Getwd()
}

func RunCmd(cmd string) (string, error) {
	command := exec.Command(cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		return "[ProcessReporterWingo.tools.RunCmd error]", err
	}
	return out.String(), nil
}

// 取两者之间
func GetMiddleStr(content string, start_str string, end_str string) string {
	n := bytes.Index([]byte(content), []byte(start_str))
	if n == -1 {
		return ""
	}
	n += len(start_str)
	m := bytes.Index([]byte(content[n:]), []byte(end_str))
	if m == -1 {
		return ""
	}
	return string(content[n : n+m])
}
