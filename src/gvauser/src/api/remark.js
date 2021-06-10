import service from '@/utils/request'

// @Tags Remark
// @Summary 创建Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "创建Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remark/createRemark [post]
export const createRemark = (data) => {
     return service({
         url: "/remark/createRemark",
         method: 'post',
         data
     })
 }


// @Tags Remark
// @Summary 删除Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "删除Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remark/deleteRemark [delete]
 export const deleteRemark = (data) => {
     return service({
         url: "/remark/deleteRemark",
         method: 'delete',
         data
     })
 }

// @Tags Remark
// @Summary 删除Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remark/deleteRemark [delete]
 export const deleteRemarkByIds = (data) => {
     return service({
         url: "/remark/deleteRemarkByIds",
         method: 'delete',
         data
     })
 }

// @Tags Remark
// @Summary 更新Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "更新Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remark/updateRemark [put]
 export const updateRemark = (data) => {
     return service({
         url: "/remark/updateRemark",
         method: 'put',
         data
     })
 }


// @Tags Remark
// @Summary 用id查询Remark
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Remark true "用id查询Remark"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remark/findRemark [get]
 export const findRemark = (params) => {
     return service({
         url: "/remark/findRemark",
         method: 'get',
         params
     })
 }


// @Tags Remark
// @Summary 分页获取Remark列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Remark列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remark/getRemarkList [get]
 export const getRemarkList = (params) => {
     return service({
         url: "/remark/getRemarkList",
         method: 'get',
         params
     })
 }