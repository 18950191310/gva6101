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

// @Tags Register
// @Summary 创建Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "创建Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /register/createRegister [post]
func CreateRegister(c *gin.Context) {
	var register model.Register
	_ = c.ShouldBindJSON(&register)
	if err := service.CreateRegister(register); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Register
// @Summary 删除Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "删除Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /register/deleteRegister [delete]
func DeleteRegister(c *gin.Context) {
	var register model.Register
	_ = c.ShouldBindJSON(&register)
	if err := service.DeleteRegister(register); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Register
// @Summary 批量删除Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /register/deleteRegisterByIds [delete]
func DeleteRegisterByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRegisterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Register
// @Summary 更新Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "更新Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /register/updateRegister [put]
func UpdateRegister(c *gin.Context) {
	var register model.Register
	_ = c.ShouldBindJSON(&register)
	if err := service.UpdateRegister(register); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Register
// @Summary 用id查询Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "用id查询Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /register/findRegister [get]
func FindRegister(c *gin.Context) {
	var register model.Register
	_ = c.ShouldBindQuery(&register)
	if err, reregister := service.GetRegister(register.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reregister": reregister}, c)
	}
}

// @Tags Register
// @Summary 分页获取Register列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RegisterSearch true "分页获取Register列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /register/getRegisterList [get]
func GetRegisterList(c *gin.Context) {
	var pageInfo request.RegisterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRegisterInfoList(pageInfo); err != nil {
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
