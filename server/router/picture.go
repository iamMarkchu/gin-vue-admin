package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPictureRouter(Router *gin.RouterGroup) {
	PictureRouter := Router.Group("picture").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		PictureRouter.POST("createPicture", v1.CreatePicture)   // 新建Picture
		PictureRouter.DELETE("deletePicture", v1.DeletePicture) // 删除Picture
		PictureRouter.DELETE("deletePictureByIds", v1.DeletePictureByIds) // 批量删除Picture
		PictureRouter.PUT("updatePicture", v1.UpdatePicture)    // 更新Picture
		PictureRouter.GET("findPicture", v1.FindPicture)        // 根据ID获取Picture
		PictureRouter.GET("getPictureList", v1.GetPictureList)  // 获取Picture列表
	}
}
