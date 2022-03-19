import service from '@/utils/request'

// @Tags Service
// @Summary 创建Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "创建Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/createService [post]
export const createService = (data) => {
  return service({
    url: '/service/createService',
    method: 'post',
    data
  })
}

// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /service/deleteService [delete]
export const deleteService = (data) => {
  return service({
    url: '/service/deleteService',
    method: 'delete',
    data
  })
}

// @Tags Service
// @Summary 删除Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /service/deleteService [delete]
export const deleteServiceByIds = (data) => {
  return service({
    url: '/service/deleteServiceByIds',
    method: 'delete',
    data
  })
}

// @Tags Service
// @Summary 更新Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Service true "更新Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /service/updateService [put]
export const updateService = (data) => {
  return service({
    url: '/service/updateService',
    method: 'put',
    data
  })
}

// @Tags Service
// @Summary 用id查询Service
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Service true "用id查询Service"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /service/findService [get]
export const findService = (params) => {
  return service({
    url: '/service/findService',
    method: 'get',
    params
  })
}

// @Tags Service
// @Summary 分页获取Service列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Service列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /service/getServiceList [get]
export const getServiceList = (params) => {
  return service({
    url: '/service/getServiceList',
    method: 'get',
    params
  })
}
