package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitGpuRouter(Router *gin.RouterGroup) {
	GpuRouter := Router.Group("gpu").Use(middleware.OperationRecord())
	{
		GpuRouter.POST("createGpu", v1.CreateGpu)   // 新建Gpu
		GpuRouter.DELETE("deleteGpu", v1.DeleteGpu) // 删除Gpu
		GpuRouter.DELETE("deleteGpuByIds", v1.DeleteGpuByIds) // 批量删除Gpu
		GpuRouter.PUT("updateGpu", v1.UpdateGpu)    // 更新Gpu
		GpuRouter.GET("findGpu", v1.FindGpu)        // 根据ID获取Gpu
		GpuRouter.GET("getGpuList", v1.GetGpuList)  // 获取Gpu列表
	}
}
