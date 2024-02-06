package requests

import (
	"log"
	"time"

	"github.com/levigross/grequests"
)

func Send(url string, data map[string]string) string {
	// 构造一个 RequestOptions
	ro := &grequests.RequestOptions{
		JSON:           data,            // JSON 字段用于设置请求体
		RequestTimeout: 5 * time.Second, // 设置 5 秒超时
	}
	// 使用 RequestOptions 发送一个 POST 请求
	resp, err := grequests.Post(url, ro)
	if err != nil {
		log.Fatalf("坏！不能发送请求: ", err)
		return err.Error()
	}
	// 打印响应体
	return resp.String()
}

// 整一个构建data的func
func BuildData() map[string]string {
	data := make(map[string]string)
	data["name"] = "TNXG"
	data["age"] = "18"
	return data
}
