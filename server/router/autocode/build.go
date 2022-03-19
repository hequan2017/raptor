package autocode

import (
	"raptor/server/api/v1"
	"raptor/server/middleware"
	"github.com/gin-gonic/gin"
)

type BuildRouter struct {
}

// InitBuildRouter 初始化 Build 路由信息
func (s *BuildRouter) InitBuildRouter(Router *gin.RouterGroup) {
	buildRouter := Router.Group("build").Use(middleware.OperationRecord())
	buildRouterWithoutRecord := Router.Group("build")
	var buildApi = v1.ApiGroupApp.AutoCodeApiGroup.BuildApi
	{
		buildRouter.POST("createBuild", buildApi.CreateBuild)   // 新建Build
		buildRouter.DELETE("deleteBuild", buildApi.DeleteBuild) // 删除Build
		buildRouter.DELETE("deleteBuildByIds", buildApi.DeleteBuildByIds) // 批量删除Build
		buildRouter.PUT("updateBuild", buildApi.UpdateBuild)    // 更新Build
	}
	{
		buildRouterWithoutRecord.GET("findBuild", buildApi.FindBuild)        // 根据ID获取Build
		buildRouterWithoutRecord.GET("getBuildList", buildApi.GetBuildList)  // 获取Build列表
	}
}
