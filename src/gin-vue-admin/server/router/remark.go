package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRemarkRouter(Router *gin.RouterGroup) {
	RemarkRouter := Router.Group("remark").Use(middleware.OperationRecord())
	{
		RemarkRouter.POST("createRemark", v1.CreateRemark)   // 新建Remark
		RemarkRouter.DELETE("deleteRemark", v1.DeleteRemark) // 删除Remark
		RemarkRouter.DELETE("deleteRemarkByIds", v1.DeleteRemarkByIds) // 批量删除Remark
		RemarkRouter.PUT("updateRemark", v1.UpdateRemark)    // 更新Remark
		RemarkRouter.GET("findRemark", v1.FindRemark)        // 根据ID获取Remark
		RemarkRouter.GET("getRemarkList", v1.GetRemarkList)  // 获取Remark列表
	}
}
