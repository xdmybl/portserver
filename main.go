package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func main() {
	InitLog()
	// 初始化 总池,因为总池是不提供设置的
	InitGlobalPool()
	InitGlobalBigJson()
	// 初始化多任务
	InitAsyncRoutine()

	logger.Debug(GlobalBigJson)
	engine := gin.New()
	engine.Use(gin.Recovery())
	api := engine.Group("/api")
	{
		api.GET("total_pool", TotalPoolVerbose)
		api.GET("pool/:id", PoolVerbose)
		api.POST("pool", AddPool)
		api.PUT("pool", UpdatePool)
		api.GET("section/:id", SectionVerbose)
		api.POST("section", AddSection)
		api.PUT("section", UpdateSection)
		api.GET("big_group/:id", BigGroupVerbose)
		api.POST("big_group", AddBigGroup)
		api.PUT("big_group", UpdateBigGroup)
		api.GET("small_group/:id", SmallGroupVerbose)
		api.POST("small_group", AddSmallGroup)
		api.PUT("small_group", UpdateSmallGroup)
		api.GET("port/:id", PortVerbose)
		api.POST("port", AddPort)
		api.PUT("port", UpdatePort)
	}
	engine.Run(":9999")
}
