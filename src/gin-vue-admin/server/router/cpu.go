package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCpuRouter(Router *gin.RouterGroup) {
	CpuRouter := Router.Group("cpu").Use(middleware.OperationRecord())
	{
		CpuRouter.POST("createCpu", v1.CreateCpu)   // 新建Cpu
		CpuRouter.DELETE("deleteCpu", v1.DeleteCpu) // 删除Cpu
		CpuRouter.DELETE("deleteCpuByIds", v1.DeleteCpuByIds) // 批量删除Cpu
		CpuRouter.PUT("updateCpu", v1.UpdateCpu)    // 更新Cpu
		CpuRouter.GET("findCpu", v1.FindCpu)        // 根据ID获取Cpu
		CpuRouter.GET("getCpuList", v1.GetCpuList)  // 获取Cpu列表
	}
}
