import service from '@/utils/request'

// @Tags Comment
// @Summary 创建Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "创建Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/createComment [post]
export const createComment = (data) => {
     return service({
         url: "/comment/createComment",
         method: 'post',
         data
     })
 }


// @Tags Comment
// @Summary 删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
 export const deleteComment = (data) => {
     return service({
         url: "/comment/deleteComment",
         method: 'delete',
         data
     })
 }

// @Tags Comment
// @Summary 删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
 export const deleteCommentByIds = (data) => {
     return service({
         url: "/comment/deleteCommentByIds",
         method: 'delete',
         data
     })
 }

// @Tags Comment
// @Summary 更新Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "更新Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comment/updateComment [put]
 export const updateComment = (data) => {
     return service({
         url: "/comment/updateComment",
         method: 'put',
         data
     })
 }


// @Tags Comment
// @Summary 用id查询Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "用id查询Comments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comments/findComments [get]
 export const findComments = (params) => {
     return service({
         url: "/comments/findComments",
         method: 'get',
         params
     })
 }


// @Tags Comments
// @Summary 分页获取Comments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Comments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comments/getCommentsList [get]
 export const getCommentsList = (params) => {
     return service({
         url: "/comments/getCommentsList",
         method: 'get',
         params
     })
 }