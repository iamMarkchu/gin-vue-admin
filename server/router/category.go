package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("category").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CategoryRouter.POST("createCategory", v1.CreateCategory)   // 新建Category
		CategoryRouter.DELETE("deleteCategory", v1.DeleteCategory) // 删除Category
		CategoryRouter.DELETE("deleteCategoryByIds", v1.DeleteCategoryByIds) // 批量删除Category
		CategoryRouter.PUT("updateCategory", v1.UpdateCategory)    // 更新Category
		CategoryRouter.GET("findCategory", v1.FindCategory)        // 根据ID获取Category
		CategoryRouter.GET("getCategoryList", v1.GetCategoryList)  // 获取Category列表
	}
}
