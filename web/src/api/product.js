import service from '@/utils/request'

// @Tags Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/createProduct [post]
export const createProduct = (data) => {
  return service({
    url: '/product/createProduct',
    method: 'post',
    data
  })
}

// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
export const deleteProduct = (data) => {
  return service({
    url: '/product/deleteProduct',
    method: 'delete',
    data
  })
}

// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
export const deleteProductByIds = (data) => {
  return service({
    url: '/product/deleteProductByIds',
    method: 'delete',
    data
  })
}

// @Tags Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /product/updateProduct [put]
export const updateProduct = (data) => {
  return service({
    url: '/product/updateProduct',
    method: 'put',
    data
  })
}

// @Tags Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /product/findProduct [get]
export const findProduct = (params) => {
  return service({
    url: '/product/findProduct',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getProductList [get]
export const getProductList = (params) => {
  return service({
    url: '/product/getProductList',
    method: 'get',
    params
  })
}
