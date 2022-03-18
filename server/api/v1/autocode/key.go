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

type KeyApi struct {
}

var keyService = service.ServiceGroupApp.AutoCodeServiceGroup.KeyService


// CreateKey 创建Key
// @Tags Key
// @Summary 创建Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Key true "创建Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /key/createKey [post]
func (keyApi *KeyApi) CreateKey(c *gin.Context) {
	var key autocode.Key
	_ = c.ShouldBindJSON(&key)
	if err := keyService.CreateKey(key); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteKey 删除Key
// @Tags Key
// @Summary 删除Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Key true "删除Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /key/deleteKey [delete]
func (keyApi *KeyApi) DeleteKey(c *gin.Context) {
	var key autocode.Key
	_ = c.ShouldBindJSON(&key)
	if err := keyService.DeleteKey(key); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteKeyByIds 批量删除Key
// @Tags Key
// @Summary 批量删除Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /key/deleteKeyByIds [delete]
func (keyApi *KeyApi) DeleteKeyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := keyService.DeleteKeyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateKey 更新Key
// @Tags Key
// @Summary 更新Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Key true "更新Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /key/updateKey [put]
func (keyApi *KeyApi) UpdateKey(c *gin.Context) {
	var key autocode.Key
	_ = c.ShouldBindJSON(&key)
	if err := keyService.UpdateKey(key); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindKey 用id查询Key
// @Tags Key
// @Summary 用id查询Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Key true "用id查询Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /key/findKey [get]
func (keyApi *KeyApi) FindKey(c *gin.Context) {
	var key autocode.Key
	_ = c.ShouldBindQuery(&key)
	if err, rekey := keyService.GetKey(key.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rekey": rekey}, c)
	}
}

// GetKeyList 分页获取Key列表
// @Tags Key
// @Summary 分页获取Key列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.KeySearch true "分页获取Key列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /key/getKeyList [get]
func (keyApi *KeyApi) GetKeyList(c *gin.Context) {
	var pageInfo autocodeReq.KeySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := keyService.GetKeyInfoList(pageInfo); err != nil {
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
