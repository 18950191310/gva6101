import service from '@/utils/request'

// @Tags Register
// @Summary 创建Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "创建Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /register/createRegister [post]
export const createRegister = (data) => {
     return service({
         url: "/register/createRegister",
         method: 'post',
         data
     })
 }


// @Tags Register
// @Summary 删除Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "删除Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /register/deleteRegister [delete]
 export const deleteRegister = (data) => {
     return service({
         url: "/register/deleteRegister",
         method: 'delete',
         data
     })
 }

// @Tags Register
// @Summary 删除Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /register/deleteRegister [delete]
 export const deleteRegisterByIds = (data) => {
     return service({
         url: "/register/deleteRegisterByIds",
         method: 'delete',
         data
     })
 }

// @Tags Register
// @Summary 更新Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "更新Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /register/updateRegister [put]
 export const updateRegister = (data) => {
     return service({
         url: "/register/updateRegister",
         method: 'put',
         data
     })
 }


// @Tags Register
// @Summary 用id查询Register
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Register true "用id查询Register"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /register/findRegister [get]
 export const findRegister = (params) => {
     return service({
         url: "/register/findRegister",
         method: 'get',
         params
     })
 }


// @Tags Register
// @Summary 分页获取Register列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Register列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /register/getRegisterList [get]
 export const getRegisterList = (params) => {
     return service({
         url: "/register/getRegisterList",
         method: 'get',
         params
     })
 }