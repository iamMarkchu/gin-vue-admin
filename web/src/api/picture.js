import service from '@/utils/request'

// @Tags Picture
// @Summary 创建Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "创建Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /picture/createPicture [post]
export const createPicture = (data) => {
     return service({
         url: "/picture/createPicture",
         method: 'post',
         data
     })
 }


// @Tags Picture
// @Summary 删除Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "删除Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /picture/deletePicture [delete]
 export const deletePicture = (data) => {
     return service({
         url: "/picture/deletePicture",
         method: 'delete',
         data
     })
 }

// @Tags Picture
// @Summary 删除Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /picture/deletePicture [delete]
 export const deletePictureByIds = (data) => {
     return service({
         url: "/picture/deletePictureByIds",
         method: 'delete',
         data
     })
 }

// @Tags Picture
// @Summary 更新Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "更新Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /picture/updatePicture [put]
 export const updatePicture = (data) => {
     return service({
         url: "/picture/updatePicture",
         method: 'put',
         data
     })
 }


// @Tags Picture
// @Summary 用id查询Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "用id查询Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /picture/findPicture [get]
 export const findPicture = (params) => {
     return service({
         url: "/picture/findPicture",
         method: 'get',
         params
     })
 }


// @Tags Picture
// @Summary 分页获取Picture列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Picture列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /picture/getPictureList [get]
 export const getPictureList = (params) => {
     return service({
         url: "/picture/getPictureList",
         method: 'get',
         params
     })
 }