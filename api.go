package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"strconv"
)

// TotalPoolVerbose
// want
// 获取总池详情
func TotalPoolVerbose(ctx *gin.Context) {
	logger.Debug("total pool: %v", GlobalBigJson.TotalPool)
	Response(ctx, httpOk, "", gin.H{"total_pool": GlobalBigJson.TotalPool}, "")
}

// ListPool
// >>> 池
// 列出
func ListPool(ctx *gin.Context) {
	Response(ctx, httpOk, "", gin.H{"pool_ls": GlobalBigJson.PoolLs}, "")
}

// PoolVerbose
// 查看详情
func PoolVerbose(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		BadResponse(ctx)
	}
	pool, err := GlobalBigJson.PoolLs.Get(id)
	if err != nil {
		BadResponse(ctx)
	}
	Response(ctx, httpOk, "", gin.H{"pool": pool}, "")
}

// AddPool
// 新增 池
func AddPool(ctx *gin.Context) {
	var bpa BasePortRangeAdd
	err := ctx.ShouldBindJSON(&bpa)
	if err != nil {
		BadResponse(ctx)
	}
	p := Pool{
		BasePorter{
			Id:          GetId(),
			Name:        bpa.Name,
			DisplayName: bpa.DisplayName,
			PortRangeLs: bpa.PortRangeLs,
			Description: bpa.Description,
			Creator:     bpa.Creator,
			CreateTime:  bpa.CreateTime,
		}}
	GlobalBigJson.PoolLs = append(GlobalBigJson.PoolLs, p)
	logger.Debug("poolLs: %v", GlobalBigJson.PoolLs)
	Response(ctx, httpOk, "", gin.H{"data": "ok"}, "")
}

// UpdatePool
// 设置 池名, 范围, 大小, 创建人, 描述
func UpdatePool(ctx *gin.Context) {
	var bpu BasePortRangeUpdate
	err := ctx.ShouldBindJSON(&bpu)
	if err != nil {
		BadResponse(ctx)
	}
	p := Pool{
		BasePorter{
			Id:          bpu.Id,
			Name:        bpu.Name,
			DisplayName: bpu.DisplayName,
			PortRangeLs: bpu.PortRangeLs,
			Description: bpu.Description,
			Creator:     bpu.Creator,
			CreateTime:  bpu.CreateTime,
		},
	}
	err = GlobalBigJson.UpdatePoolLsFromId(bpu.Id, p)
	if err != nil {
		BadResponse(ctx)
	}
	fmt.Printf("poolLs: %v", GlobalBigJson.PoolLs)
	Response(ctx, httpOk, "", gin.H{"data": p}, "")
}

// ListSection
// >>> 段
// 列出
func ListSection(ctx *gin.Context) {
	Response(ctx, httpOk, "", gin.H{"section_ls": GlobalBigJson.SectionLs}, "")
}

// SectionVerbose
// 查看详情
func SectionVerbose(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		BadResponse(ctx)
	}
	section, err := GlobalBigJson.SectionLs.Get(id)
	if err != nil {
		BadResponse(ctx)
	}
	Response(ctx, httpOk, "", gin.H{"data": section}, "")
}

// AddSection
// 新增 段
func AddSection(ctx *gin.Context) {
	var spa SubPortRangeAdd
	err := ctx.ShouldBindJSON(&spa)
	if err != nil {
		BadResponse(ctx)
	}
	s := Section{
		BasePorter: BasePorter{
			Id:          GetId(),
			Name:        spa.Name,
			DisplayName: spa.DisplayName,
			PortRangeLs: spa.PortRangeLs,
			Description: spa.Description,
			Creator:     spa.Creator,
			CreateTime:  spa.CreateTime,
		},
		SuperId: spa.SuperId,
	}
	GlobalBigJson.SectionLs = append(GlobalBigJson.SectionLs, s)
	logger.Debug("sectionLs: %v", GlobalBigJson.SectionLs)
	Response(ctx, httpOk, "", gin.H{"data": "ok"}, "")
}

// UpdateSection
// 设置 段名, 范围, 大小, 创建人, 描述
func UpdateSection(ctx *gin.Context) {
	var spu SubPortRangeUpdate
	err := ctx.ShouldBindJSON(&spu)
	if err != nil {
		BadResponse(ctx)
	}
	s := Section{
		BasePorter: BasePorter{
			Id:          spu.Id,
			Name:        spu.Name,
			DisplayName: spu.DisplayName,
			PortRangeLs: spu.PortRangeLs,
			Description: spu.Description,
			Creator:     spu.Creator,
			CreateTime:  spu.CreateTime,
		},
		SuperId: spu.SuperId,
	}
	err = GlobalBigJson.UpdateSectionLsFromId(spu.Id, s)
	if err != nil {
		BadResponse(ctx)
	}
	fmt.Printf("sectionLs: %v", GlobalBigJson.PoolLs)
	Response(ctx, httpOk, "", gin.H{"data": s}, "")
}

// ListBigGroup
// >>> 大组
// 列出
func ListBigGroup(ctx *gin.Context) {
	Response(ctx, httpOk, "", gin.H{"big_group_ls": GlobalBigJson.BigGroupLs}, "")
}

// BigGroupVerbose
// 查看详情
func BigGroupVerbose(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		BadResponse(ctx)
	}
	bigGroup, err := GlobalBigJson.BigGroupLs.Get(id)
	if err != nil {
		BadResponse(ctx)
	}
	Response(ctx, httpOk, "", gin.H{"data": bigGroup}, "")
}

// AddBigGroup
// 新增 大组
func AddBigGroup(ctx *gin.Context) {
	var spa SubPortRangeAdd
	err := ctx.ShouldBindJSON(&spa)
	if err != nil {
		BadResponse(ctx)
	}
	s := BigGroup{
		BasePorter: BasePorter{
			Id:          GetId(),
			Name:        spa.Name,
			DisplayName: spa.DisplayName,
			PortRangeLs: spa.PortRangeLs,
			Description: spa.Description,
			Creator:     spa.Creator,
			CreateTime:  spa.CreateTime,
		},
		SuperId: spa.SuperId,
	}
	GlobalBigJson.BigGroupLs = append(GlobalBigJson.BigGroupLs, s)
	logger.Debug("big_group_ls: %v", GlobalBigJson.BigGroupLs)
	Response(ctx, httpOk, "", gin.H{"data": "ok"}, "")
}

// UpdateBigGroup
// 设置 大组名, 范围, 大小, 创建人, 描述
func UpdateBigGroup(ctx *gin.Context) {
	var spu SubPortRangeUpdate
	err := ctx.ShouldBindJSON(&spu)
	if err != nil {
		BadResponse(ctx)
	}
	s := BigGroup{
		BasePorter: BasePorter{
			Id:          spu.Id,
			Name:        spu.Name,
			DisplayName: spu.DisplayName,
			PortRangeLs: spu.PortRangeLs,
			Description: spu.Description,
			Creator:     spu.Creator,
			CreateTime:  spu.CreateTime,
		},
		SuperId: spu.SuperId,
	}
	err = GlobalBigJson.UpdateBigGroupLsFromId(spu.Id, s)
	if err != nil {
		BadResponse(ctx)
	}
	fmt.Printf("big_group_ls: %v", GlobalBigJson.BigGroupLs)
	Response(ctx, httpOk, "", gin.H{"data": s}, "")
}

// ListSmallGroup 列出
func ListSmallGroup(ctx *gin.Context) {
	Response(ctx, httpOk, "", gin.H{"small_group_ls": GlobalBigJson.SmallGroupLs}, "")
}

// SmallGroupVerbose
// >>> 小组
// 查看详情
func SmallGroupVerbose(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		BadResponse(ctx)
	}
	smallGroup, err := GlobalBigJson.SmallGroupLs.Get(id)
	if err != nil {
		BadResponse(ctx)
	}
	Response(ctx, httpOk, "", gin.H{"data": smallGroup}, "")
}

// AddSmallGroup 新增 小组
func AddSmallGroup(ctx *gin.Context) {
	var spa SubPortRangeAdd
	err := ctx.ShouldBindJSON(&spa)
	if err != nil {
		BadResponse(ctx)
	}
	s := SmallGroup{
		BasePorter: BasePorter{
			Id:          GetId(),
			Name:        spa.Name,
			DisplayName: spa.DisplayName,
			PortRangeLs: spa.PortRangeLs,
			Description: spa.Description,
			Creator:     spa.Creator,
			CreateTime:  spa.CreateTime,
		},
		SuperId: spa.SuperId,
	}
	GlobalBigJson.SmallGroupLs = append(GlobalBigJson.SmallGroupLs, s)
	logger.Debug("small_group_ls: %v", GlobalBigJson.SmallGroupLs)
	Response(ctx, httpOk, "", gin.H{"data": "ok"}, "")
}

// UpdateSmallGroup 设置 小组名, 范围, 大小, 创建人, 描述
func UpdateSmallGroup(ctx *gin.Context) {
	var spu SubPortRangeUpdate
	err := ctx.ShouldBindJSON(&spu)
	if err != nil {
		BadResponse(ctx)
	}
	s := SmallGroup{
		BasePorter: BasePorter{
			Id:          spu.Id,
			Name:        spu.Name,
			DisplayName: spu.DisplayName,
			PortRangeLs: spu.PortRangeLs,
			Description: spu.Description,
			Creator:     spu.Creator,
			CreateTime:  spu.CreateTime,
		},
		SuperId: spu.SuperId,
	}
	err = GlobalBigJson.UpdateSmallGroupLsFromId(spu.Id, s)
	if err != nil {
		BadResponse(ctx)
	}
	fmt.Printf("small_group_ls: %v", GlobalBigJson.SmallGroupLs)
	Response(ctx, httpOk, "", gin.H{"data": s}, "")
}

// ListPort 列出
func ListPort(ctx *gin.Context) {
	Response(ctx, httpOk, "", gin.H{"port_ls": GlobalBigJson.PortLs}, "")
}

// PortVerbose
// >>> 端口
// 查看详情
func PortVerbose(ctx *gin.Context) {
	stringId := ctx.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		BadResponse(ctx)
	}
	port, err := GlobalBigJson.PortLs.Get(id)
	if err != nil {
		BadResponse(ctx)
	}
	Response(ctx, httpOk, "", gin.H{"data": port}, "")
}

// AddPort 新增 端口
func AddPort(ctx *gin.Context) {
	var pa PortAdd
	err := ctx.ShouldBindJSON(&pa)
	if err != nil {
		BadResponse(ctx)
	}
	p := Port{
		Id:          GetId(),
		Name:        pa.Name,
		SuperId:     pa.SuperId,
		DisplayName: pa.DisplayName,
		PortValue:   pa.PortValue,
		Description: pa.Description,
		Creator:     pa.Creator,
		CreateTime:  pa.CreateTime,
	}
	GlobalBigJson.PortLs = append(GlobalBigJson.PortLs, p)
	logger.Debug("port: %v", GlobalBigJson.SmallGroupLs)
	Response(ctx, httpOk, "", gin.H{"data": "ok"}, "")
}

// UpdatePort 设置 端口, 范围, 大小, 创建人, 描述
func UpdatePort(ctx *gin.Context) {
	var pa PortUpdate
	err := ctx.ShouldBindJSON(&pa)
	if err != nil {
		BadResponse(ctx)
	}

	p := Port{
		Id:          pa.Id,
		Name:        pa.Name,
		SuperId:     pa.SuperId,
		DisplayName: pa.DisplayName,
		PortValue:   pa.PortValue,
		Description: pa.Description,
		Creator:     pa.Creator,
		CreateTime:  pa.CreateTime,
	}
	err = GlobalBigJson.UpdatePortLsFromId(pa.Id, p)
	if err != nil {
		BadResponse(ctx)
	}
	fmt.Printf("port_ls: %v", GlobalBigJson.PortLs)
	Response(ctx, httpOk, "", gin.H{"data": p}, "")
}
