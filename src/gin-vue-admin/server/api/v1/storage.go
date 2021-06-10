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

// @Tags Storage
// @Summary 创建Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "创建Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /storage/createStorage [post]
func CreateStorage(c *gin.Context) {
	var storage model.Storage
	_ = c.ShouldBindJSON(&storage)
	if err := service.CreateStorage(storage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Storage
// @Summary 删除Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "删除Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /storage/deleteStorage [delete]
func DeleteStorage(c *gin.Context) {
	var storage model.Storage
	_ = c.ShouldBindJSON(&storage)
	if err := service.DeleteStorage(storage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Storage
// @Summary 批量删除Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /storage/deleteStorageByIds [delete]
func DeleteStorageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteStorageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Storage
// @Summary 更新Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "更新Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /storage/updateStorage [put]
func UpdateStorage(c *gin.Context) {
	var storage model.Storage
	_ = c.ShouldBindJSON(&storage)
	if err := service.UpdateStorage(storage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Storage
// @Summary 用id查询Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "用id查询Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /storage/findStorage [get]
func FindStorage(c *gin.Context) {
	var storage model.Storage
	_ = c.ShouldBindQuery(&storage)
	if err, restorage := service.GetStorage(storage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restorage": restorage}, c)
	}
}

// @Tags Storage
// @Summary 分页获取Storage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.StorageSearch true "分页获取Storage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /storage/getStorageList [get]
func GetStorageList(c *gin.Context) {
	var pageInfo request.StorageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetStorageInfoList(pageInfo); err != nil {
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
