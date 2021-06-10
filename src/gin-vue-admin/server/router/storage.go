package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitStorageRouter(Router *gin.RouterGroup) {
	StorageRouter := Router.Group("storage").Use(middleware.OperationRecord())
	{
		StorageRouter.POST("createStorage", v1.CreateStorage)   // 新建Storage
		StorageRouter.DELETE("deleteStorage", v1.DeleteStorage) // 删除Storage
		StorageRouter.DELETE("deleteStorageByIds", v1.DeleteStorageByIds) // 批量删除Storage
		StorageRouter.PUT("updateStorage", v1.UpdateStorage)    // 更新Storage
		StorageRouter.GET("findStorage", v1.FindStorage)        // 根据ID获取Storage
		StorageRouter.GET("getStorageList", v1.GetStorageList)  // 获取Storage列表
	}
}
