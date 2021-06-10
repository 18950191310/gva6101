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

// @Tags Recolist
// @Summary 创建Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "创建Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recolist/createRecolist [post]
func CreateRecolist(c *gin.Context) {
	var recolist model.Recolist
	_ = c.ShouldBindJSON(&recolist)
	if err := service.CreateRecolist(recolist); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Recolist
// @Summary 删除Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "删除Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recolist/deleteRecolist [delete]
func DeleteRecolist(c *gin.Context) {
	var recolist model.Recolist
	_ = c.ShouldBindJSON(&recolist)
	if err := service.DeleteRecolist(recolist); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Recolist
// @Summary 批量删除Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /recolist/deleteRecolistByIds [delete]
func DeleteRecolistByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRecolistByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Recolist
// @Summary 更新Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "更新Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recolist/updateRecolist [put]
func UpdateRecolist(c *gin.Context) {
	var recolist model.Recolist
	_ = c.ShouldBindJSON(&recolist)
	if err := service.UpdateRecolist(recolist); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Recolist
// @Summary 用id查询Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "用id查询Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recolist/findRecolist [get]
func FindRecolist(c *gin.Context) {
	var recolist model.Recolist
	_ = c.ShouldBindQuery(&recolist)
	if err, rerecolist := service.GetRecolist(recolist.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerecolist": rerecolist}, c)
	}
}
// @Tags Recolist
// @Summary 分页获取Recolist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RecolistSearch true "分页获取Recolist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recolist/getRecolistList [get]
func GetRecolistList(c *gin.Context) {
	var pageInfo request.RecolistSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRecolistInfoList(pageInfo); err != nil {
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
