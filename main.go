package main

import (
	"log"
	"strings"
	"time"

	Config "github.com/TNXG/ProcessReporterWingo/config"
	Core "github.com/TNXG/ProcessReporterWingo/core"
	Requests "github.com/TNXG/ProcessReporterWingo/core/requests"
	Setup "github.com/TNXG/ProcessReporterWingo/core/setup"
)

// 当前程序版本
var Version = "0.0.2"

// 从配置文件中读取配置信息
var cfg = Config.LoadConfig()

// 服务器的端点
var endpoint = cfg.ServerConfig.Endpoint

// 服务器的令牌
var token = cfg.ServerConfig.Token

// 报告时间间隔（秒）
var reportTime = cfg.ServerConfig.ReportTime

// 获取项目初始化信息
var setupstatus = Setup.Setup(Version)

// report 函数用于报告当前前台窗口的进程信息
func report() {
	// 提示更新
	if setupstatus {
		log.Printf("程序有新版本！请更新！")
	}
	// 获取当前前台窗口的进程名
	processName, _ := Core.GetWindowInfo()
	// 处理一下进程名
	processName = strings.TrimSuffix(processName, ".exe")
	processName = Core.Replacer(processName)
	// 创建一个空的媒体更新map
	var mediaUpdate map[string]string
	// 获取media信息
	Title, Artist, SourceAppName := Core.GetSmtcInfo()
	if Title != "" {
		mediaUpdate = Requests.BuildMediaUpdate(Title, Artist, SourceAppName)
	} else {
		mediaUpdate = map[string]string{}
	}
	// 构建数据map，包含时间戳、进程名、媒体更新和token四个键
	updateData := Requests.BuildData(processName, mediaUpdate, token)
	// 向指定的endpoint发送POST请求，请求的数据是updateData
	Requests.Report(updateData, endpoint)
}

func main() {
	// 创建一个新的定时器，每隔reportTime秒就会触发一次
	ticker := time.NewTicker(time.Duration(reportTime) * time.Second)
	// 确保在程序结束时停止定时器
	defer ticker.Stop()

	// 每当定时器触发时，就调用report函数
	for range ticker.C {
		report()
	}
}
