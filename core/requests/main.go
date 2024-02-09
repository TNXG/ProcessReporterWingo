package Requests

import (
	"encoding/json"
	"log"
	"time"
	"unsafe"

	"github.com/levigross/grequests"
)

// StringToCharPtr 函数将一个字符串转换为字符指针
func StringToCharPtr(str string) *uint16 {
	chars := append([]byte(str), 0)
	return (*uint16)(unsafe.Pointer(&chars[0]))
}

// BuildMediaUpdate 函数用于构建一个媒体更新的map，包含"title"和"artist"两个键
func BuildMediaUpdate(title, artist, SourceAppName string) map[string]string {
	return map[string]string{
		"title":         title,
		"artist":        artist,
		"SourceAppName": SourceAppName,
	}
}

// BuildData 函数用于构建一个数据map，包含时间戳、进程名、媒体更新和token四个键
// 如果媒体更新中的"title"为空，则不会包含"media"键
func BuildData(processName string, mediaUpdate map[string]string, token string) map[string]interface{} {
	timestamp := int(time.Now().Unix())

	var updateData map[string]interface{}
	if mediaUpdate["title"] != "" {
		updateData = map[string]interface{}{
			"timestamp": timestamp,
			"process":   processName,
			"media":     mediaUpdate,
			"key":       token,
		}
	} else {
		updateData = map[string]interface{}{
			"timestamp": timestamp,
			"process":   processName,
			"key":       token,
		}
	}

	return updateData
}

// Report 函数用于向指定的endpoint发送POST请求，请求的数据是updateData，请求头包含"Content-Type"和"User-Agent"两个键
// 如果请求发送失败或者响应解析失败，会打印错误信息
// 如果请求和响应都成功，会打印API的响应
func Report(updateData map[string]interface{}, endpoint string) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; TokaiTeio) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82 iykrzu/114.514",
	}

	resp, err := grequests.Post(endpoint, &grequests.RequestOptions{
		Headers:        headers,
		UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; TokaiTeio) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82 iykrzu/114.514",
		JSON:           updateData,
		RequestTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("发送请求时出错：%v\n", err)
		return
	}
	defer resp.Close()
	var response map[string]interface{}
	if err := json.Unmarshal(resp.Bytes(), &response); err != nil {
		log.Printf("解析响应时出错：%v\n", err)
		return
	}
	log.Printf("API 响应：%+v\n", response)
}
