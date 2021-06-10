import service from '@/utils/request'

// @Tags Gpu
// @Summary 创建Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "创建Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpu/createGpu [post]
export const createGpu = (data) => {
     return service({
         url: "/gpu/createGpu",
         method: 'post',
         data
     })
 }


// @Tags Gpu
// @Summary 删除Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "删除Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpu/deleteGpu [delete]
 export const deleteGpu = (data) => {
     return service({
         url: "/gpu/deleteGpu",
         method: 'delete',
         data
     })
 }

// @Tags Gpu
// @Summary 删除Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gpu/deleteGpu [delete]
 export const deleteGpuByIds = (data) => {
     return service({
         url: "/gpu/deleteGpuByIds",
         method: 'delete',
         data
     })
 }

// @Tags Gpu
// @Summary 更新Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "更新Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gpu/updateGpu [put]
 export const updateGpu = (data) => {
     return service({
         url: "/gpu/updateGpu",
         method: 'put',
         data
     })
 }


// @Tags Gpu
// @Summary 用id查询Gpu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gpu true "用id查询Gpu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gpu/findGpu [get]
 export const findGpu = (params) => {
     return service({
         url: "/gpu/findGpu",
         method: 'get',
         params
     })
 }


// @Tags Gpu
// @Summary 分页获取Gpu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Gpu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gpu/getGpuList [get]
 export const getGpuList = (params) => {
     return service({
         url: "/gpu/getGpuList",
         method: 'get',
         params
     })
 }