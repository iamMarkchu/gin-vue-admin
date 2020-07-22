package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Picture
// @Summary 创建Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "创建Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /picture/createPicture [post]
func CreatePicture(c *gin.Context) {
	var picture model.Picture
	_ = c.ShouldBindJSON(&picture)
	err := service.CreatePicture(picture)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Picture
// @Summary 删除Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "删除Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /picture/deletePicture [delete]
func DeletePicture(c *gin.Context) {
	var picture model.Picture
	_ = c.ShouldBindJSON(&picture)
	err := service.DeletePicture(picture)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Picture
// @Summary 批量删除Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /picture/deletePictureByIds [delete]
func DeletePictureByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeletePictureByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Picture
// @Summary 更新Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "更新Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /picture/updatePicture [put]
func UpdatePicture(c *gin.Context) {
	var picture model.Picture
	_ = c.ShouldBindJSON(&picture)
	err := service.UpdatePicture(&picture)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Picture
// @Summary 用id查询Picture
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Picture true "用id查询Picture"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /picture/findPicture [get]
func FindPicture(c *gin.Context) {
	var picture model.Picture
	_ = c.ShouldBindQuery(&picture)
	err, repicture := service.GetPicture(picture.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"repicture": repicture}, c)
	}
}

// @Tags Picture
// @Summary 分页获取Picture列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PictureSearch true "分页获取Picture列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /picture/getPictureList [get]
func GetPictureList(c *gin.Context) {
	var pageInfo request.PictureSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetPictureInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
