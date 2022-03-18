package autocode

import (
	"github.com/gin-gonic/gin"
	"raptor/server/api/v1"
	"raptor/server/middleware"
)

type KeyRouter struct {
}

// InitKeyRouter 初始化 Key 路由信息
func (s *KeyRouter) InitKeyRouter(Router *gin.RouterGroup) {
	keyRouter := Router.Group("key").Use(middleware.OperationRecord())
	keyRouterWithoutRecord := Router.Group("key")
	var keyApi = v1.ApiGroupApp.AutoCodeApiGroup.KeyApi
	{
		keyRouter.POST("createKey", keyApi.CreateKey)   // 新建Key
		keyRouter.DELETE("deleteKey", keyApi.DeleteKey) // 删除Key
		keyRouter.DELETE("deleteKeyByIds", keyApi.DeleteKeyByIds) // 批量删除Key
		keyRouter.PUT("updateKey", keyApi.UpdateKey)    // 更新Key
	}
	{
		keyRouterWithoutRecord.GET("findKey", keyApi.FindKey)        // 根据ID获取Key
		keyRouterWithoutRecord.GET("getKeyList", keyApi.GetKeyList)  // 获取Key列表
	}
}
