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

// @Tags Remark
// @Summary 创建Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "创建Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remark/createRemark [post]
func CreateRemark(c *gin.Context) {
	var remark model.Remark
	_ = c.ShouldBindJSON(&remark)
	if err := service.CreateRemark(remark); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Remark
// @Summary 删除Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "删除Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remark/deleteRemark [delete]
func DeleteRemark(c *gin.Context) {
	var remark model.Remark
	_ = c.ShouldBindJSON(&remark)
	if err := service.DeleteRemark(remark); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Remark
// @Summary 批量删除Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /remark/deleteRemarkByIds [delete]
func DeleteRemarkByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRemarkByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Remark
// @Summary 更新Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "更新Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remark/updateRemark [put]
func UpdateRemark(c *gin.Context) {
	var remark model.Remark
	_ = c.ShouldBindJSON(&remark)
	if err := service.UpdateRemark(remark); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Remark
// @Summary 用id查询Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "用id查询Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remark/findRemark [get]
func FindRemark(c *gin.Context) {
	var remark model.Remark
	_ = c.ShouldBindQuery(&remark)
	if err, reremark := service.GetRemark(remark.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reremark": reremark}, c)
	}
}

// @Tags Remark
// @Summary 分页获取Remark列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RemarkSearch true "分页获取Remark列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remark/getRemarkList [get]
func GetRemarkList(c *gin.Context) {
	var pageInfo request.RemarkSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRemarkInfoList(pageInfo); err != nil {
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
