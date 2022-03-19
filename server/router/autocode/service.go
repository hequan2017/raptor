package autocode

import (
	"raptor/server/api/v1"
	"raptor/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceRouter struct {
}

// InitServiceRouter 初始化 Service 路由信息
func (s *ServiceRouter) InitServiceRouter(Router *gin.RouterGroup) {
	serviceRouter := Router.Group("service").Use(middleware.OperationRecord())
	serviceRouterWithoutRecord := Router.Group("service")
	var serviceApi = v1.ApiGroupApp.AutoCodeApiGroup.ServiceApi
	{
		serviceRouter.POST("createService", serviceApi.CreateService)   // 新建Service
		serviceRouter.DELETE("deleteService", serviceApi.DeleteService) // 删除Service
		serviceRouter.DELETE("deleteServiceByIds", serviceApi.DeleteServiceByIds) // 批量删除Service
		serviceRouter.PUT("updateService", serviceApi.UpdateService)    // 更新Service
	}
	{
		serviceRouterWithoutRecord.GET("findService", serviceApi.FindService)        // 根据ID获取Service
		serviceRouterWithoutRecord.GET("getServiceList", serviceApi.GetServiceList)  // 获取Service列表
	}
}
