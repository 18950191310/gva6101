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

// @Tags CuserDetail
// @Summary 创建CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "创建CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cuserDetail/createCuserDetail [post]
func CreateCuserDetail(c *gin.Context) {
	var cuserDetail model.CuserDetail
	_ = c.ShouldBindJSON(&cuserDetail)
	if err := service.CreateCuserDetail(cuserDetail); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags CuserDetail
// @Summary 删除CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "删除CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cuserDetail/deleteCuserDetail [delete]
func DeleteCuserDetail(c *gin.Context) {
	var cuserDetail model.CuserDetail
	_ = c.ShouldBindJSON(&cuserDetail)
	if err := service.DeleteCuserDetail(cuserDetail); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags CuserDetail
// @Summary 批量删除CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cuserDetail/deleteCuserDetailByIds [delete]
func DeleteCuserDetailByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCuserDetailByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags CuserDetail
// @Summary 更新CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "更新CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cuserDetail/updateCuserDetail [put]
func UpdateCuserDetail(c *gin.Context) {
	var cuserDetail model.CuserDetail
	_ = c.ShouldBindJSON(&cuserDetail)
	if err := service.UpdateCuserDetail(cuserDetail); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags CuserDetail
// @Summary 用id查询CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "用id查询CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cuserDetail/findCuserDetail [get]
func FindCuserDetail(c *gin.Context) {
	var cuserDetail model.CuserDetail
	_ = c.ShouldBindQuery(&cuserDetail)
	if err, recuserDetail := service.GetCuserDetail(cuserDetail.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recuserDetail": recuserDetail}, c)
	}
}

// @Tags CuserDetail
// @Summary 分页获取CuserDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CuserDetailSearch true "分页获取CuserDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cuserDetail/getCuserDetailList [get]
func GetCuserDetailList(c *gin.Context) {
	var pageInfo request.CuserDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCuserDetailInfoList(pageInfo); err != nil {
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
