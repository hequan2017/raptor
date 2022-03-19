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

type ServiceApi struct {
}

var serviceService = service.ServiceGroupApp.AutoCodeServiceGroup.ServiceService


// CreateService 创建Service
// @Tags Service
// @Summary 创建Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Service true "创建Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/createService [post]
func (serviceApi *ServiceApi) CreateService(c *gin.Context) {
	var service autocode.Service
	_ = c.ShouldBindJSON(&service)
	if err := serviceService.CreateService(service); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteService 删除Service
// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Service true "删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /service/deleteService [delete]
func (serviceApi *ServiceApi) DeleteService(c *gin.Context) {
	var service autocode.Service
	_ = c.ShouldBindJSON(&service)
	if err := serviceService.DeleteService(service); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServiceByIds 批量删除Service
// @Tags Service
// @Summary 批量删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /service/deleteServiceByIds [delete]
func (serviceApi *ServiceApi) DeleteServiceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := serviceService.DeleteServiceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateService 更新Service
// @Tags Service
// @Summary 更新Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Service true "更新Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /service/updateService [put]
func (serviceApi *ServiceApi) UpdateService(c *gin.Context) {
	var service autocode.Service
	_ = c.ShouldBindJSON(&service)
	if err := serviceService.UpdateService(service); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindService 用id查询Service
// @Tags Service
// @Summary 用id查询Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Service true "用id查询Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /service/findService [get]
func (serviceApi *ServiceApi) FindService(c *gin.Context) {
	var service autocode.Service
	_ = c.ShouldBindQuery(&service)
	if err, reservice := serviceService.GetService(service.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reservice": reservice}, c)
	}
}

// GetServiceList 分页获取Service列表
// @Tags Service
// @Summary 分页获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ServiceSearch true "分页获取Service列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/getServiceList [get]
func (serviceApi *ServiceApi) GetServiceList(c *gin.Context) {
	var pageInfo autocodeReq.ServiceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := serviceService.GetServiceInfoList(pageInfo); err != nil {
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
