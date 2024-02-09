package setup

import (
	"log"
	"os"
	"strings"

	"github.com/TNXG/ProcessReporterWingo/tools"
	"github.com/levigross/grequests"
)

func Setup(Version string) bool {
	workdir, _ := tools.Getwd()
	// 查询本地是否存在version.txt
	_, err := os.Stat(workdir + "/version.txt")
	var nowversion string
	if err != nil {
		file, _ := os.Create(workdir + "/version.txt")
		defer file.Close()
		file.WriteString(Version)
		nowversion = Version
	} else {
		// 如果存在则读取文件内容
		file, _ := os.Open(workdir + "/version.txt")
		defer file.Close()
		buf := make([]byte, 100)
		n, _ := file.Read(buf)
		// 如果文件内容不是当前版本号，则更新文件内容
		nowversion = strings.TrimSpace(string(buf[:n]))
	}
	// 从githubapi获取最新的Release版本号
	resp, _ := grequests.Get("https://tnxg-proxy.deno.dev/https://api.github.com/repos/TNXG/ProcessReporterWingo/releases/latest", nil)

	var result map[string]interface{}
	if err := resp.JSON(&result); err != nil {
		// 打印错误信息并返回
		log.Printf("Error parsing JSON:")
		return false
	}
	latestversion := result["tag_name"].(string)
	// 如果不相同说明需要更新
	var UpdateRequired bool
	if latestversion != nowversion {
		UpdateRequired = true
	} else {
		UpdateRequired = false
	}

	return UpdateRequired
}
