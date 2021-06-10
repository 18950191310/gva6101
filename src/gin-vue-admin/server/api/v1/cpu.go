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

// @Tags Cpu
// @Summary 创建Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "创建Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpu/createCpu [post]
func CreateCpu(c *gin.Context) {
	var cpu model.Cpu
	_ = c.ShouldBindJSON(&cpu)
	if err := service.CreateCpu(cpu); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Cpu
// @Summary 删除Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "删除Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpu/deleteCpu [delete]
func DeleteCpu(c *gin.Context) {
	var cpu model.Cpu
	_ = c.ShouldBindJSON(&cpu)
	if err := service.DeleteCpu(cpu); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Cpu
// @Summary 批量删除Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cpu/deleteCpuByIds [delete]
func DeleteCpuByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCpuByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Cpu
// @Summary 更新Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "更新Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpu/updateCpu [put]
func UpdateCpu(c *gin.Context) {
	var cpu model.Cpu
	_ = c.ShouldBindJSON(&cpu)
	if err := service.UpdateCpu(cpu); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Cpu
// @Summary 用id查询Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "用id查询Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpu/findCpu [get]
func FindCpu(c *gin.Context) {
	var cpu model.Cpu
	_ = c.ShouldBindQuery(&cpu)
	if err, recpu := service.GetCpu(cpu.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recpu": recpu}, c)
	}
}

// @Tags Cpu
// @Summary 分页获取Cpu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CpuSearch true "分页获取Cpu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpu/getCpuList [get]
func GetCpuList(c *gin.Context) {
	var pageInfo request.CpuSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCpuInfoList(pageInfo); err != nil {
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
