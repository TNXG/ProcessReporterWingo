package setup

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/TNXG/ProcessReporterWingo/tools"
	"github.com/levigross/grequests"
)

// Setup 函数用于检查是否需要更新应用程序
func Setup(Version string) bool {
	workdir, _ := tools.Getwd() // 获取当前工作目录
	// 查询本地是否存在version.txt
	_, err := os.Stat(workdir + "/version.txt")
	var nowversion string
	if err != nil {
		// 如果文件不存在，创建一个新的version.txt文件，并写入当前版本号
		file, _ := os.Create(workdir + "/version.txt")
		defer file.Close()
		file.WriteString(Version)
		nowversion = Version
	} else {
		// 如果文件存在则读取文件内容
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
		// 如果解析JSON出错，打印错误信息并返回false
		log.Printf("Error parsing JSON:")
		return false
	}
	latestversion := result["tag_name"].(string)
	// 获取 /core/GetSmtcInfo.exe是否存在，不存在则下载
	// 先创建一个core目录
	_, err = os.Stat(workdir + "/core")
	if err != nil {
		os.Mkdir(workdir+"/core", 0755)
	}
	// 再下载
	_, err = os.Stat(workdir + "/core/GetSmtcInfo.exe")
	if err != nil {
		log.Printf("GetSmtcInfo.exe 不存在，正在下载...")
		// 下载GetSmtcInfo.exe
		resp, _ := grequests.Get("https://tnxg-proxy.deno.dev/https://github.com/TNXG/ProcessReporterWingo/raw/master/core/GetSmtcInfo.exe", nil)
		if !resp.Ok {
			log.Printf("下载 GetSmtcInfo.exe 出错, 状态码 %d", resp.StatusCode)
		} else {
			// 将下载的文件写入到本地
			err = ioutil.WriteFile(workdir+"/core/GetSmtcInfo.exe", resp.Bytes(), 0755)
			if err != nil {
				log.Printf("GetSmtcInfo.exe写入错误: %s", err)
			}
		}
	}
	// 如果最新版本号与当前版本号不同，说明需要更新，返回true
	if latestversion != nowversion {
		return true
	} else {
		// 否则，不需要更新，返回false
		return false
	}

}
