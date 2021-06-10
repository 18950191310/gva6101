import service from '@/utils/request'

// @Tags Storage
// @Summary 创建Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "创建Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /storage/createStorage [post]
export const createStorage = (data) => {
     return service({
         url: "/storage/createStorage",
         method: 'post',
         data
     })
 }


// @Tags Storage
// @Summary 删除Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "删除Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /storage/deleteStorage [delete]
 export const deleteStorage = (data) => {
     return service({
         url: "/storage/deleteStorage",
         method: 'delete',
         data
     })
 }

// @Tags Storage
// @Summary 删除Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /storage/deleteStorage [delete]
 export const deleteStorageByIds = (data) => {
     return service({
         url: "/storage/deleteStorageByIds",
         method: 'delete',
         data
     })
 }

// @Tags Storage
// @Summary 更新Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "更新Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /storage/updateStorage [put]
 export const updateStorage = (data) => {
     return service({
         url: "/storage/updateStorage",
         method: 'put',
         data
     })
 }


// @Tags Storage
// @Summary 用id查询Storage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Storage true "用id查询Storage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /storage/findStorage [get]
 export const findStorage = (params) => {
     return service({
         url: "/storage/findStorage",
         method: 'get',
         params
     })
 }


// @Tags Storage
// @Summary 分页获取Storage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Storage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /storage/getStorageList [get]
 export const getStorageList = (params) => {
     return service({
         url: "/storage/getStorageList",
         method: 'get',
         params
     })
 }