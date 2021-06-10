import service from '@/utils/request'

// @Tags Cpu
// @Summary 创建Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "创建Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpu/createCpu [post]
export const createCpu = (data) => {
     return service({
         url: "/cpu/createCpu",
         method: 'post',
         data
     })
 }


// @Tags Cpu
// @Summary 删除Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "删除Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpu/deleteCpu [delete]
 export const deleteCpu = (data) => {
     return service({
         url: "/cpu/deleteCpu",
         method: 'delete',
         data
     })
 }

// @Tags Cpu
// @Summary 删除Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpu/deleteCpu [delete]
 export const deleteCpuByIds = (data) => {
     return service({
         url: "/cpu/deleteCpuByIds",
         method: 'delete',
         data
     })
 }

// @Tags Cpu
// @Summary 更新Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "更新Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpu/updateCpu [put]
 export const updateCpu = (data) => {
     return service({
         url: "/cpu/updateCpu",
         method: 'put',
         data
     })
 }


// @Tags Cpu
// @Summary 用id查询Cpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cpu true "用id查询Cpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpu/findCpu [get]
 export const findCpu = (params) => {
     return service({
         url: "/cpu/findCpu",
         method: 'get',
         params
     })
 }


// @Tags Cpu
// @Summary 分页获取Cpu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Cpu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpu/getCpuList [get]
 export const getCpuList = (params) => {
     return service({
         url: "/cpu/getCpuList",
         method: 'get',
         params
     })
 }