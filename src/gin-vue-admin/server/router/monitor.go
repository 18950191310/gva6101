package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMonitorRouter(Router *gin.RouterGroup) {
	MonitorRouter := Router.Group("monitor").Use(middleware.OperationRecord())
	{
		MonitorRouter.POST("createMonitor", v1.CreateMonitor)   // 新建Monitor
		MonitorRouter.DELETE("deleteMonitor", v1.DeleteMonitor) // 删除Monitor
		MonitorRouter.DELETE("deleteMonitorByIds", v1.DeleteMonitorByIds) // 批量删除Monitor
		MonitorRouter.PUT("updateMonitor", v1.UpdateMonitor)    // 更新Monitor
		MonitorRouter.GET("findMonitor", v1.FindMonitor)        // 根据ID获取Monitor
		MonitorRouter.GET("getMonitorList", v1.GetMonitorList)  // 获取Monitor列表
	}
}
