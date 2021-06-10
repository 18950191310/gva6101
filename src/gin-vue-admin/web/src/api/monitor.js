import service from '@/utils/request'

// @Tags Monitor
// @Summary 创建Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "创建Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitor/createMonitor [post]
export const createMonitor = (data) => {
     return service({
         url: "/monitor/createMonitor",
         method: 'post',
         data
     })
 }


// @Tags Monitor
// @Summary 删除Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "删除Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitor/deleteMonitor [delete]
 export const deleteMonitor = (data) => {
     return service({
         url: "/monitor/deleteMonitor",
         method: 'delete',
         data
     })
 }

// @Tags Monitor
// @Summary 删除Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monitor/deleteMonitor [delete]
 export const deleteMonitorByIds = (data) => {
     return service({
         url: "/monitor/deleteMonitorByIds",
         method: 'delete',
         data
     })
 }

// @Tags Monitor
// @Summary 更新Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "更新Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monitor/updateMonitor [put]
 export const updateMonitor = (data) => {
     return service({
         url: "/monitor/updateMonitor",
         method: 'put',
         data
     })
 }


// @Tags Monitor
// @Summary 用id查询Monitor
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Monitor true "用id查询Monitor"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monitor/findMonitor [get]
 export const findMonitor = (params) => {
     return service({
         url: "/monitor/findMonitor",
         method: 'get',
         params
     })
 }


// @Tags Monitor
// @Summary 分页获取Monitor列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Monitor列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monitor/getMonitorList [get]
 export const getMonitorList = (params) => {
     return service({
         url: "/monitor/getMonitorList",
         method: 'get',
         params
     })
 }