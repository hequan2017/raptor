package autocode

import (
	"raptor/server/global"
    "raptor/server/model/autocode"
    "raptor/server/model/common/request"
    autocodeReq "raptor/server/model/autocode/request"
    "raptor/server/model/common/response"
    "raptor/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BuildApi struct {
}

var buildService = service.ServiceGroupApp.AutoCodeServiceGroup.BuildService


// CreateBuild 创建Build
// @Tags Build
// @Summary 创建Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Build true "创建Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /build/createBuild [post]
func (buildApi *BuildApi) CreateBuild(c *gin.Context) {
	var build autocode.Build
	_ = c.ShouldBindJSON(&build)
	if err := buildService.CreateBuild(build); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBuild 删除Build
// @Tags Build
// @Summary 删除Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Build true "删除Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /build/deleteBuild [delete]
func (buildApi *BuildApi) DeleteBuild(c *gin.Context) {
	var build autocode.Build
	_ = c.ShouldBindJSON(&build)
	if err := buildService.DeleteBuild(build); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBuildByIds 批量删除Build
// @Tags Build
// @Summary 批量删除Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /build/deleteBuildByIds [delete]
func (buildApi *BuildApi) DeleteBuildByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := buildService.DeleteBuildByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBuild 更新Build
// @Tags Build
// @Summary 更新Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Build true "更新Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /build/updateBuild [put]
func (buildApi *BuildApi) UpdateBuild(c *gin.Context) {
	var build autocode.Build
	_ = c.ShouldBindJSON(&build)
	if err := buildService.UpdateBuild(build); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBuild 用id查询Build
// @Tags Build
// @Summary 用id查询Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Build true "用id查询Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /build/findBuild [get]
func (buildApi *BuildApi) FindBuild(c *gin.Context) {
	var build autocode.Build
	_ = c.ShouldBindQuery(&build)
	if err, rebuild := buildService.GetBuild(build.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebuild": rebuild}, c)
	}
}

// GetBuildList 分页获取Build列表
// @Tags Build
// @Summary 分页获取Build列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BuildSearch true "分页获取Build列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /build/getBuildList [get]
func (buildApi *BuildApi) GetBuildList(c *gin.Context) {
	var pageInfo autocodeReq.BuildSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := buildService.GetBuildInfoList(pageInfo); err != nil {
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
