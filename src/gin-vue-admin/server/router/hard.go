package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHardRouter(Router *gin.RouterGroup) {
	HardRouter := Router.Group("hard").Use(middleware.OperationRecord())
	{
		HardRouter.POST("createHard", v1.CreateHard)   // 新建Hard
		HardRouter.DELETE("deleteHard", v1.DeleteHard) // 删除Hard
		HardRouter.DELETE("deleteHardByIds", v1.DeleteHardByIds) // 批量删除Hard
		HardRouter.PUT("updateHard", v1.UpdateHard)    // 更新Hard
		HardRouter.GET("findHard", v1.FindHard)        // 根据ID获取Hard
		HardRouter.GET("getHardList", v1.GetHardList)  // 获取Hard列表
	}
}
