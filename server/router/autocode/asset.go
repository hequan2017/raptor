package autocode

import (
	"github.com/gin-gonic/gin"
	"raptor/server/api/v1"
	"raptor/server/middleware"
)

type AssetRouter struct {
}

// InitAssetRouter 初始化 Asset 路由信息
func (s *AssetRouter) InitAssetRouter(Router *gin.RouterGroup) {
	assetRouter := Router.Group("asset").Use(middleware.OperationRecord())
	assetRouterWithoutRecord := Router.Group("asset")
	var assetApi = v1.ApiGroupApp.AutoCodeApiGroup.AssetApi
	{
		assetRouter.POST("createAsset", assetApi.CreateAsset)   // 新建Asset
		assetRouter.DELETE("deleteAsset", assetApi.DeleteAsset) // 删除Asset
		assetRouter.DELETE("deleteAssetByIds", assetApi.DeleteAssetByIds) // 批量删除Asset
		assetRouter.PUT("updateAsset", assetApi.UpdateAsset)    // 更新Asset
	}
	{
		assetRouterWithoutRecord.GET("findAsset", assetApi.FindAsset)        // 根据ID获取Asset
		assetRouterWithoutRecord.GET("getAssetList", assetApi.GetAssetList)  // 获取Asset列表
		assetRouterWithoutRecord.GET("syncAsset", assetApi.SyncAsset)        // 同步更新
	}
}
