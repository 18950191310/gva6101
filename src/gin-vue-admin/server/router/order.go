package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("order").Use(middleware.OperationRecord())
	{
		OrderRouter.POST("createOrder", v1.CreateOrder)   // 新建Order
		OrderRouter.DELETE("deleteOrder", v1.DeleteOrder) // 删除Order
		OrderRouter.DELETE("deleteOrderByIds", v1.DeleteOrderByIds) // 批量删除Order
		OrderRouter.PUT("updateOrder", v1.UpdateOrder)    // 更新Order
		OrderRouter.GET("findOrder", v1.FindOrder)        // 根据ID获取Order
		OrderRouter.GET("getOrderList", v1.GetOrderList)  // 获取Order列表
	}
}
