package autocode

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"raptor/server/global"
	"raptor/server/model/autocode"
	autocodeReq "raptor/server/model/autocode/request"
	"raptor/server/model/common/request"
	"raptor/server/model/common/response"
	"raptor/server/service"
	"raptor/server/task"
)

type AssetApi struct {
}

var assetService = service.ServiceGroupApp.AutoCodeServiceGroup.AssetService


// CreateAsset 创建Asset
// @Tags Asset
// @Summary 创建Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Asset true "创建Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset/createAsset [post]
func (assetApi *AssetApi) CreateAsset(c *gin.Context) {
	var asset autocode.Asset
	_ = c.ShouldBindJSON(&asset)
	if err := assetService.CreateAsset(asset); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAsset 删除Asset
// @Tags Asset
// @Summary 删除Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Asset true "删除Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset/deleteAsset [delete]
func (assetApi *AssetApi) DeleteAsset(c *gin.Context) {
	var asset autocode.Asset
	_ = c.ShouldBindJSON(&asset)
	if err := assetService.DeleteAsset(asset); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAssetByIds 批量删除Asset
// @Tags Asset
// @Summary 批量删除Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /asset/deleteAssetByIds [delete]
func (assetApi *AssetApi) DeleteAssetByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := assetService.DeleteAssetByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAsset 更新Asset
// @Tags Asset
// @Summary 更新Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Asset true "更新Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asset/updateAsset [put]
func (assetApi *AssetApi) UpdateAsset(c *gin.Context) {
	var asset autocode.Asset
	_ = c.ShouldBindJSON(&asset)
	if err := assetService.UpdateAsset(asset); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAsset 用id查询Asset
// @Tags Asset
// @Summary 用id查询Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Asset true "用id查询Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset/findAsset [get]
func (assetApi *AssetApi) FindAsset(c *gin.Context) {
	var asset autocode.Asset
	_ = c.ShouldBindQuery(&asset)
	if err, reasset := assetService.GetAsset(asset.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reasset": reasset}, c)
	}
}

// GetAssetList 分页获取Asset列表
// @Tags Asset
// @Summary 分页获取Asset列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AssetSearch true "分页获取Asset列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset/getAssetList [get]
func (assetApi *AssetApi) GetAssetList(c *gin.Context) {
	var pageInfo autocodeReq.AssetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := assetService.GetAssetInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// SyncAsset  同步
// @Tags Asset
// @Summary 同步
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Asset true 同步
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset/SyncAsset [get]
func (assetApi *AssetApi) SyncAsset(c *gin.Context) {
	go task.AssetUpdate()
	response.OkWithData(gin.H{"reasset": "正在同步中"}, c)
}