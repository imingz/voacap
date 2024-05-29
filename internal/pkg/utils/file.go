package utils

import (
	"runtime"
	"strings"
)

// GetFilePath 根据 runtime.GOOS() 获取当前系统类型，进而获取文件执行路径
func GetFilePath(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}
	return "/mnt/" + strings.ToLower(path[:1]) + "/" + strings.ReplaceAll(path[3:], "\\", "/")
}
