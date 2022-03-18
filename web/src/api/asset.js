import service from '@/utils/request'

// @Tags Asset
// @Summary 创建Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset true "创建Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset/createAsset [post]
export const createAsset = (data) => {
  return service({
    url: '/asset/createAsset',
    method: 'post',
    data
  })
}

// @Tags Asset
// @Summary 删除Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset true "删除Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset/deleteAsset [delete]
export const deleteAsset = (data) => {
  return service({
    url: '/asset/deleteAsset',
    method: 'delete',
    data
  })
}

// @Tags Asset
// @Summary 删除Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /asset/deleteAsset [delete]
export const deleteAssetByIds = (data) => {
  return service({
    url: '/asset/deleteAssetByIds',
    method: 'delete',
    data
  })
}

// @Tags Asset
// @Summary 更新Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Asset true "更新Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /asset/updateAsset [put]
export const updateAsset = (data) => {
  return service({
    url: '/asset/updateAsset',
    method: 'put',
    data
  })
}

// @Tags Asset
// @Summary 用id查询Asset
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Asset true "用id查询Asset"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /asset/findAsset [get]
export const findAsset = (params) => {
  return service({
    url: '/asset/findAsset',
    method: 'get',
    params
  })
}

// @Tags Asset
// @Summary 分页获取Asset列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Asset列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /asset/getAssetList [get]
export const getAssetList = (params) => {
  return service({
    url: '/asset/getAssetList',
    method: 'get',
    params
  })
}
