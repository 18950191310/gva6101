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

// @Tags Hard
// @Summary 创建Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "创建Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hard/createHard [post]
func CreateHard(c *gin.Context) {
	var hard model.Hard
	_ = c.ShouldBindJSON(&hard)
	if err := service.CreateHard(hard); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Hard
// @Summary 删除Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "删除Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hard/deleteHard [delete]
func DeleteHard(c *gin.Context) {
	var hard model.Hard
	_ = c.ShouldBindJSON(&hard)
	if err := service.DeleteHard(hard); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hard
// @Summary 批量删除Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hard/deleteHardByIds [delete]
func DeleteHardByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteHardByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Hard
// @Summary 更新Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "更新Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hard/updateHard [put]
func UpdateHard(c *gin.Context) {
	var hard model.Hard
	_ = c.ShouldBindJSON(&hard)
	if err := service.UpdateHard(hard); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Hard
// @Summary 用id查询Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "用id查询Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hard/findHard [get]
func FindHard(c *gin.Context) {
	var hard model.Hard
	_ = c.ShouldBindQuery(&hard)
	if err, rehard := service.GetHard(hard.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehard": rehard}, c)
	}
}

// @Tags Hard
// @Summary 分页获取Hard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HardSearch true "分页获取Hard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hard/getHardList [get]
func GetHardList(c *gin.Context) {
	var pageInfo request.HardSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHardInfoList(pageInfo); err != nil {
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
