package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRegisterRouter(Router *gin.RouterGroup) {
	RegisterRouter := Router.Group("register").Use(middleware.OperationRecord())
	{
		RegisterRouter.POST("createRegister", v1.CreateRegister)   // 新建Register
		RegisterRouter.DELETE("deleteRegister", v1.DeleteRegister) // 删除Register
		RegisterRouter.DELETE("deleteRegisterByIds", v1.DeleteRegisterByIds) // 批量删除Register
		RegisterRouter.PUT("updateRegister", v1.UpdateRegister)    // 更新Register
		RegisterRouter.GET("findRegister", v1.FindRegister)        // 根据ID获取Register
		RegisterRouter.GET("getRegisterList", v1.GetRegisterList)  // 获取Register列表
	}
}
