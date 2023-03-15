package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

func Response(ctx *gin.Context, code string, msg string, data gin.H, info string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
		"info":    info,
	})
}

func BadResponse(ctx *gin.Context) {
	Response(ctx, httpNok, "", nil, "")
}

func InitGlobalPool() {
	//GlobalTotalPool = &TotalPool{
	//	BasePorter: &BasePorter{
	//		Name:        "total_pool",
	//		DisplayName: "总池",
	//		PortRangeLs: []PortRange{*TotalPoolRange},
	//		Description: "总的端口池",
	//		Creator:     "xiaoy",
	//		CreateTime:  "2023/3/6",
	//	},
	//}
	GlobalTotalPool.Init()
	//t := &TotalPool{}
	//GlobalTotalPool = t.Init()
	//GlobalTotalPool = t
	logger.Debug(GlobalTotalPool)
}

func InitAsyncRoutine() {
	go SaveJson()
}

// SaveJson 保存 Global
func SaveJson() {
	for {
		logger.Trace("save file")
		err := GlobalBigJson.SaveToFile()
		if err != nil {
			logger.Error("save file error: %v\n", err)
		}
		time.Sleep(5 * time.Second)
	}
}

// GetId
// 获取一个 Id
func GetId() int {
	GlobalBigJson.Lock()
	GlobalBigJson.MaxId++
	GlobalBigJson.Unlock()
	return GlobalBigJson.MaxId
}
