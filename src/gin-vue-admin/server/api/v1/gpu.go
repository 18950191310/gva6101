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

// @Tags Gpu
// @Summary 创建Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "创建Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpu/createGpu [post]
func CreateGpu(c *gin.Context) {
	var gpu model.Gpu
	_ = c.ShouldBindJSON(&gpu)
	if err := service.CreateGpu(gpu); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Gpu
// @Summary 删除Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "删除Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpu/deleteGpu [delete]
func DeleteGpu(c *gin.Context) {
	var gpu model.Gpu
	_ = c.ShouldBindJSON(&gpu)
	if err := service.DeleteGpu(gpu); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Gpu
// @Summary 批量删除Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /gpu/deleteGpuByIds [delete]
func DeleteGpuByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteGpuByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Gpu
// @Summary 更新Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "更新Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gpu/updateGpu [put]
func UpdateGpu(c *gin.Context) {
	var gpu model.Gpu
	_ = c.ShouldBindJSON(&gpu)
	if err := service.UpdateGpu(gpu); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Gpu
// @Summary 用id查询Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "用id查询Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gpu/findGpu [get]
func FindGpu(c *gin.Context) {
	var gpu model.Gpu
	_ = c.ShouldBindQuery(&gpu)
	if err, regpu := service.GetGpu(gpu.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regpu": regpu}, c)
	}
}

// @Tags Gpu
// @Summary 分页获取Gpu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GpuSearch true "分页获取Gpu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpu/getGpuList [get]
func GetGpuList(c *gin.Context) {
	var pageInfo request.GpuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetGpuInfoList(pageInfo); err != nil {
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
