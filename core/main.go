package Core

import (
	"io/ioutil"
	"log"
	"syscall"
	"unsafe"

	"github.com/TNXG/ProcessReporterWingo/tools"
	"gopkg.in/yaml.v2"
)

// Config 结构体用于存储从YAML配置文件中读取的配置信息
type Config struct {
	Server struct {
		Endpoint   string `yaml:"endpoint"`   // 服务器的端点
		Token      string `yaml:"token"`      // 令牌
		ReportTime int    `yaml:"ReportTime"` // 报告时间
	} `yaml:"server"` // 服务器配置
}

// ReadConf 函数用于读取和解析YAML配置文件
func ReadConf() Config {
	// 读取文件
	workdir, _ := tools.Getwd()
	data, err := ioutil.ReadFile(workdir + "/config.yml")
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}
	// 将YAML反序列化为结构体
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("无法解析文件: %v", err)
	}
	// 直接返回
	return conf
}

// 定义一些变量，用于调用Windows API
var (
	user32DLL                    = syscall.MustLoadDLL("User32.dll")                             // 加载User32.dll
	procGetForegroundWindow      = user32DLL.MustFindProc("GetForegroundWindow")                 // 获取GetForegroundWindow函数的地址
	procGetWindowTextW           = user32DLL.MustFindProc("GetWindowTextW")                      // 获取GetWindowTextW函数的地址
	procGetWindowThreadProcessId = user32DLL.MustFindProc("GetWindowThreadProcessId")            // 获取GetWindowThreadProcessId函数的地址
	procOpenProcess              = syscall.NewLazyDLL("kernel32.dll").NewProc("OpenProcess")     // 获取OpenProcess函数的地址
	procGetModuleBaseNameW       = syscall.NewLazyDLL("psapi.dll").NewProc("GetModuleBaseNameW") // 获取GetModuleBaseNameW函数的地址
)

// StringToCharPtr 函数用于将字符串转换为字符指针
func StringToCharPtr(str string) *uint16 {
	chars := append([]byte(str), 0)             // 将字符串转换为字节切片，并在末尾添加一个零字节
	return (*uint16)(unsafe.Pointer(&chars[0])) // 将字节切片的首地址转换为*uint16类型并返回
}

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
