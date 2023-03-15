package main

import (
	"github.com/wonderivan/logger"
	"os"
)

// InitGlobalBigJson
// 副作用
func InitGlobalBigJson() {
	// 加载 bigJson
	bigJson, err := LoadJson()
	if err != nil {
		logger.Debug("err: LoadJson error")
		// 第一次读不到文件
		os.Exit(-1)
	}
	GlobalBigJson = bigJson
}
