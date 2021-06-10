package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRecolistRouter(Router *gin.RouterGroup) {
	RecolistRouter := Router.Group("recolist").Use(middleware.OperationRecord())
	{
		RecolistRouter.POST("createRecolist", v1.CreateRecolist)   // 新建Recolist
		RecolistRouter.DELETE("deleteRecolist", v1.DeleteRecolist) // 删除Recolist
		RecolistRouter.DELETE("deleteRecolistByIds", v1.DeleteRecolistByIds) // 批量删除Recolist
		RecolistRouter.PUT("updateRecolist", v1.UpdateRecolist)    // 更新Recolist
		RecolistRouter.GET("findRecolist", v1.FindRecolist)        // 根据ID获取Recolist
		RecolistRouter.GET("getRecolistList", v1.GetRecolistList)  // 获取Recolist列表
	}
}
