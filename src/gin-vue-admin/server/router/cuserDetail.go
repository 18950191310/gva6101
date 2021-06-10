package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCuserDetailRouter(Router *gin.RouterGroup) {
	CuserDetailRouter := Router.Group("cuserDetail").Use(middleware.OperationRecord())
	{
		CuserDetailRouter.POST("createCuserDetail", v1.CreateCuserDetail)   // 新建CuserDetail
		CuserDetailRouter.DELETE("deleteCuserDetail", v1.DeleteCuserDetail) // 删除CuserDetail
		CuserDetailRouter.DELETE("deleteCuserDetailByIds", v1.DeleteCuserDetailByIds) // 批量删除CuserDetail
		CuserDetailRouter.PUT("updateCuserDetail", v1.UpdateCuserDetail)    // 更新CuserDetail
		CuserDetailRouter.GET("findCuserDetail", v1.FindCuserDetail)        // 根据ID获取CuserDetail
		CuserDetailRouter.GET("getCuserDetailList", v1.GetCuserDetailList)  // 获取CuserDetail列表
	}
}
