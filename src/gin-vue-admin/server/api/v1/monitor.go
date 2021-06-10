package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Monitor
// @Summary 创建Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "创建Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitor/createMonitor [post]
func CreateMonitor(c *gin.Context) {
	var monitor model.Monitor
	_ = c.ShouldBindJSON(&monitor)
	if err := service.CreateMonitor(monitor); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Monitor
// @Summary 删除Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "删除Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitor/deleteMonitor [delete]
func DeleteMonitor(c *gin.Context) {
	var monitor model.Monitor
	_ = c.ShouldBindJSON(&monitor)
	if err := service.DeleteMonitor(monitor); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Monitor
// @Summary 批量删除Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /monitor/deleteMonitorByIds [delete]
func DeleteMonitorByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMonitorByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Monitor
// @Summary 更新Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "更新Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monitor/updateMonitor [put]
func UpdateMonitor(c *gin.Context) {
	var monitor model.Monitor
	_ = c.ShouldBindJSON(&monitor)
	if err := service.UpdateMonitor(monitor); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Monitor
// @Summary 用id查询Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "用id查询Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monitor/findMonitor [get]
func FindMonitor(c *gin.Context) {
	var monitor model.Monitor
	_ = c.ShouldBindQuery(&monitor)
	if err, remonitor := service.GetMonitor(monitor.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remonitor": remonitor}, c)
	}
}

// @Tags Monitor
// @Summary 分页获取Monitor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MonitorSearch true "分页获取Monitor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitor/getMonitorList [get]
func GetMonitorList(c *gin.Context) {
	var pageInfo request.MonitorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMonitorInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
