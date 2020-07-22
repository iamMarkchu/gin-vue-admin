package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCategory
// @description   create a Category
// @param     category               model.Category
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCategory(category model.Category) (err error) {
	category.Status = 1
	category.UserId = 1
	err = global.GVA_DB.Create(&category).Error
	return err
}

// @title    DeleteCategory
// @description   delete a Category
// @auth                     （2020/04/05  20:22）
// @param     category               model.Category
// @return                    error

func DeleteCategory(category model.Category) (err error) {
	err = global.GVA_DB.Delete(category).Error
	return err
}

// @title    DeleteCategoryByIds
// @description   delete Categorys
// @auth                     （2020/04/05  20:22）
// @param     category               model.Category
// @return                    error

func DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Category{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdateCategory
// @description   update a Category
// @param     category          *model.Category
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCategory(category *model.Category) (err error) {
	err = global.GVA_DB.Save(category).Error
	return err
}

// @title    GetCategory
// @description   get the info of a Category
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Category        Category

func GetCategory(id uint) (err error, category model.Category) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// @title    GetCategoryInfoList
// @description   get Category list by pagination, 分页获取Category
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCategoryInfoList(info request.CategorySearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Category{})
    var categorys []model.Category
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CatName != "" {
        db = db.Where("`cat_name` LIKE ?","%"+ info.CatName+"%")
    }
    if info.Description != "" {
        db = db.Where("`description` LIKE ?","%"+ info.Description+"%")
    }
    if info.Pid != 0 {
        db = db.Where("`pid` = ?",info.Pid)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return err, categorys, total
}
