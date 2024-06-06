//go:generate goversioninfo
package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

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

var RSMediaTitle, RSMediaArtist, RSSourceAppName = "", "", ""

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
	// 构建数据map，包含时间戳、进程名、媒体更新和token四个键
	updateData := Requests.BuildData(processName, getMediaMessage(), token)
	// 向指定的endpoint发送POST请求，请求的数据是updateData
	Requests.Report(updateData, endpoint)
}

func main() {
	// 在启动时启动ReportServer
	go ReportServer()

	// 创建一个新的定时器，每隔reportTime秒就会触发一次
	ticker := time.NewTicker(time.Duration(reportTime) * time.Second)
	// 确保在程序结束时停止定时器
	defer ticker.Stop()

	// 每当定时器触发时，就调用report函数
	for range ticker.C {
		report()
	}
}

func getMediaMessage() map[string]string {
	// 获取media信息
	Title, Artist, SourceAppName := Core.GetSmtcInfo()

	if RSMediaTitle != "" {
		return Requests.BuildMediaUpdate(RSMediaTitle, RSMediaArtist, RSSourceAppName)
	} else if Title != "" {
		return Requests.BuildMediaUpdate(Title, Artist, SourceAppName)
	}
	return map[string]string{}
}

func ReportServer() {
	// 创建Echo实例
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "这里是ProcessReporterWingo的上报api",
		})
	})
	// 定义一个GET请求的处理器函数
	e.POST("/api/report/media", func(c echo.Context) error {
		// 获取请求参数
		Status := c.Param("status")
		if Status == "start" {
			RSMediaTitle, RSMediaArtist, RSSourceAppName = c.Param("title"), c.Param("artist"), c.Param("SourceAppName")
		} else if Status == "close" {
			RSMediaTitle, RSMediaArtist, RSSourceAppName = "", "", ""
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Report received and processed.",
		})
	})

	e.GET("/api/report", func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{
			"message": "Invalid request method. Only POST is allowed.",
		})
	})

	e.Start(":4138")
}
