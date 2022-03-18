import service from '@/utils/request'

// @Tags Key
// @Summary 创建Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "创建Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /key/createKey [post]
export const createKey = (data) => {
  return service({
    url: '/key/createKey',
    method: 'post',
    data
  })
}

// @Tags Key
// @Summary 删除Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "删除Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /key/deleteKey [delete]
export const deleteKey = (data) => {
  return service({
    url: '/key/deleteKey',
    method: 'delete',
    data
  })
}

// @Tags Key
// @Summary 删除Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /key/deleteKey [delete]
export const deleteKeyByIds = (data) => {
  return service({
    url: '/key/deleteKeyByIds',
    method: 'delete',
    data
  })
}

// @Tags Key
// @Summary 更新Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Key true "更新Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /key/updateKey [put]
export const updateKey = (data) => {
  return service({
    url: '/key/updateKey',
    method: 'put',
    data
  })
}

// @Tags Key
// @Summary 用id查询Key
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Key true "用id查询Key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /key/findKey [get]
export const findKey = (params) => {
  return service({
    url: '/key/findKey',
    method: 'get',
    params
  })
}

// @Tags Key
// @Summary 分页获取Key列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Key列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /key/getKeyList [get]
export const getKeyList = (params) => {
  return service({
    url: '/key/getKeyList',
    method: 'get',
    params
  })
}
