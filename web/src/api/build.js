import service from '@/utils/request'

// @Tags Build
// @Summary 创建Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Build true "创建Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /build/createBuild [post]
export const createBuild = (data) => {
  return service({
    url: '/build/createBuild',
    method: 'post',
    data
  })
}

// @Tags Build
// @Summary 删除Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Build true "删除Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /build/deleteBuild [delete]
export const deleteBuild = (data) => {
  return service({
    url: '/build/deleteBuild',
    method: 'delete',
    data
  })
}

// @Tags Build
// @Summary 删除Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /build/deleteBuild [delete]
export const deleteBuildByIds = (data) => {
  return service({
    url: '/build/deleteBuildByIds',
    method: 'delete',
    data
  })
}

// @Tags Build
// @Summary 更新Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Build true "更新Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /build/updateBuild [put]
export const updateBuild = (data) => {
  return service({
    url: '/build/updateBuild',
    method: 'put',
    data
  })
}

// @Tags Build
// @Summary 用id查询Build
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Build true "用id查询Build"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /build/findBuild [get]
export const findBuild = (params) => {
  return service({
    url: '/build/findBuild',
    method: 'get',
    params
  })
}

// @Tags Build
// @Summary 分页获取Build列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Build列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /build/getBuildList [get]
export const getBuildList = (params) => {
  return service({
    url: '/build/getBuildList',
    method: 'get',
    params
  })
}
