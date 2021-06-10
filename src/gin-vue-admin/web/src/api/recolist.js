import service from '@/utils/request'

// @Tags Recolist
// @Summary 创建Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "创建Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recolist/createRecolist [post]
export const createRecolist = (data) => {
     return service({
         url: "/recolist/createRecolist",
         method: 'post',
         data
     })
 }


// @Tags Recolist
// @Summary 删除Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "删除Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recolist/deleteRecolist [delete]
 export const deleteRecolist = (data) => {
     return service({
         url: "/recolist/deleteRecolist",
         method: 'delete',
         data
     })
 }

// @Tags Recolist
// @Summary 删除Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recolist/deleteRecolist [delete]
 export const deleteRecolistByIds = (data) => {
     return service({
         url: "/recolist/deleteRecolistByIds",
         method: 'delete',
         data
     })
 }

// @Tags Recolist
// @Summary 更新Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "更新Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recolist/updateRecolist [put]
 export const updateRecolist = (data) => {
     return service({
         url: "/recolist/updateRecolist",
         method: 'put',
         data
     })
 }


// @Tags Recolist
// @Summary 用id查询Recolist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recolist true "用id查询Recolist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recolist/findRecolist [get]
 export const findRecolist = (params) => {
     return service({
         url: "/recolist/findRecolist",
         method: 'get',
         params
     })
 }


// @Tags Recolist
// @Summary 分页获取Recolist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Recolist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recolist/getRecolistList [get]
 export const getRecolistList = (params) => {
     return service({
         url: "/recolist/getRecolistList",
         method: 'get',
         params
     })
 }