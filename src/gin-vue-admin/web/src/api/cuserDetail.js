import service from '@/utils/request'

// @Tags CuserDetail
// @Summary 创建CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "创建CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cuserDetail/createCuserDetail [post]
export const createCuserDetail = (data) => {
     return service({
         url: "/cuserDetail/createCuserDetail",
         method: 'post',
         data
     })
 }


// @Tags CuserDetail
// @Summary 删除CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "删除CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cuserDetail/deleteCuserDetail [delete]
 export const deleteCuserDetail = (data) => {
     return service({
         url: "/cuserDetail/deleteCuserDetail",
         method: 'delete',
         data
     })
 }

// @Tags CuserDetail
// @Summary 删除CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cuserDetail/deleteCuserDetail [delete]
 export const deleteCuserDetailByIds = (data) => {
     return service({
         url: "/cuserDetail/deleteCuserDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags CuserDetail
// @Summary 更新CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "更新CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cuserDetail/updateCuserDetail [put]
 export const updateCuserDetail = (data) => {
     return service({
         url: "/cuserDetail/updateCuserDetail",
         method: 'put',
         data
     })
 }


// @Tags CuserDetail
// @Summary 用id查询CuserDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CuserDetail true "用id查询CuserDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cuserDetail/findCuserDetail [get]
 export const findCuserDetail = (params) => {
     return service({
         url: "/cuserDetail/findCuserDetail",
         method: 'get',
         params
     })
 }


// @Tags CuserDetail
// @Summary 分页获取CuserDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CuserDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cuserDetail/getCuserDetailList [get]
 export const getCuserDetailList = (params) => {
     return service({
         url: "/cuserDetail/getCuserDetailList",
         method: 'get',
         params
     })
 }