package util

import "time"

// 获取当前时间字符串
func GetLocalNowTimeStr() string {
	localTime := time.Now().Format("2006-01-02 15:04:05")
	return localTime
}
