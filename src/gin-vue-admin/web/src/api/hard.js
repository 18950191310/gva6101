import service from '@/utils/request'

// @Tags Hard
// @Summary 创建Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "创建Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hard/createHard [post]
export const createHard = (data) => {
     return service({
         url: "/hard/createHard",
         method: 'post',
         data
     })
 }


// @Tags Hard
// @Summary 删除Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "删除Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hard/deleteHard [delete]
 export const deleteHard = (data) => {
     return service({
         url: "/hard/deleteHard",
         method: 'delete',
         data
     })
 }

// @Tags Hard
// @Summary 删除Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hard/deleteHard [delete]
 export const deleteHardByIds = (data) => {
     return service({
         url: "/hard/deleteHardByIds",
         method: 'delete',
         data
     })
 }

// @Tags Hard
// @Summary 更新Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "更新Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hard/updateHard [put]
 export const updateHard = (data) => {
     return service({
         url: "/hard/updateHard",
         method: 'put',
         data
     })
 }


// @Tags Hard
// @Summary 用id查询Hard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hard true "用id查询Hard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hard/findHard [get]
 export const findHard = (params) => {
     return service({
         url: "/hard/findHard",
         method: 'get',
         params
     })
 }


// @Tags Hard
// @Summary 分页获取Hard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Hard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hard/getHardList [get]
 export const getHardList = (params) => {
     return service({
         url: "/hard/getHardList",
         method: 'get',
         params
     })
 }