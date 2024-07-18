package Core

import (
	"syscall"
	"unsafe"

	Config "github.com/TNXG/ProcessReporterWingo/config"
	"github.com/TNXG/ProcessReporterWingo/tools"
)

// 定义一些变量
var (
	user32DLL                    = syscall.MustLoadDLL("User32.dll")                             // 加载User32.dll
	procGetForegroundWindow      = user32DLL.MustFindProc("GetForegroundWindow")                 // 获取GetForegroundWindow函数的地址
	procGetWindowTextW           = user32DLL.MustFindProc("GetWindowTextW")                      // 获取GetWindowTextW函数的地址
	procGetWindowThreadProcessId = user32DLL.MustFindProc("GetWindowThreadProcessId")            // 获取GetWindowThreadProcessId函数的地址
	procOpenProcess              = syscall.NewLazyDLL("kernel32.dll").NewProc("OpenProcess")     // 获取OpenProcess函数的地址
	procGetModuleBaseNameW       = syscall.NewLazyDLL("psapi.dll").NewProc("GetModuleBaseNameW") // 获取GetModuleBaseNameW函数的地址
)

// GetWindowInfo 函数用于获取当前活动窗口的信息
// 返回两个字符串，第一个是进程名，第二个是窗口标题
func GetWindowInfo() (string, string) {
	hWnd, _, _ := procGetForegroundWindow.Call()                                                 // 获取当前活动窗口的句柄
	windowTitle := make([]uint16, 255)                                                           // 创建一个长度为255的uint16切片，用于存储窗口的标题
	procGetWindowTextW.Call(hWnd, uintptr(unsafe.Pointer(&windowTitle[0])), 255)                 // 获取窗口的标题
	var processID uint32                                                                         // 创建一个uint32变量，用于存储进程的ID
	procGetWindowThreadProcessId.Call(hWnd, uintptr(unsafe.Pointer(&processID)))                 // 获取进程的ID
	processHandle, _, _ := procOpenProcess.Call(0x0400|0x0010, 0, uintptr(processID))            // 获取进程的句柄
	processName := make([]uint16, 255)                                                           // 创建一个长度为255的uint16切片，用于存储进程的名字
	procGetModuleBaseNameW.Call(processHandle, 0, uintptr(unsafe.Pointer(&processName[0])), 255) // 获取进程的名字
	return syscall.UTF16ToString(processName), syscall.UTF16ToString(windowTitle)                // 返回进程的名字和窗口的标题
}

// 通过使用C#编写的程序获取媒体信息
func GetSmtcInfo() (string, string, string) {
	workdir, _ := tools.Getwd()
	// 运行程序获取输出
	result, _ := tools.RunCmd(workdir + "/core/GetSmtcInfo.exe")
	Title := tools.GetMiddleStr(result, "[ProcessReporterWingo.GetSmtcInfo.Title]", "[/ProcessReporterWingo.GetSmtcInfo.Title]")
	Artist := tools.GetMiddleStr(result, "[ProcessReporterWingo.GetSmtcInfo.Artist]", "[/ProcessReporterWingo.GetSmtcInfo.Artist]")
	SourceAppName := tools.GetMiddleStr(result, "[ProcessReporterWingo.GetSmtcInfo.SourceAppName]", "[/ProcessReporterWingo.GetSmtcInfo.SourceAppName]")
	return Title, Artist, SourceAppName
}

// 替换文本内容
func Replacer(processName string) string {
	cfg := Config.LoadConfig()
	for _, rule := range cfg.Rules {
		if processName == rule.MatchApplication {
			return rule.Replace.Application
		}
	}
	return processName
}
