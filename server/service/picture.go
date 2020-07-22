package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreatePicture
// @description   create a Picture
// @param     picture               model.Picture
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreatePicture(picture model.Picture) (err error) {
	picture.Status = 1
	picture.UserId = 1
	err = global.GVA_DB.Create(&picture).Error
	return err
}

// @title    DeletePicture
// @description   delete a Picture
// @auth                     （2020/04/05  20:22）
// @param     picture               model.Picture
// @return                    error

func DeletePicture(picture model.Picture) (err error) {
	err = global.GVA_DB.Delete(picture).Error
	return err
}

// @title    DeletePictureByIds
// @description   delete Pictures
// @auth                     （2020/04/05  20:22）
// @param     picture               model.Picture
// @return                    error

func DeletePictureByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Picture{},"id in (?)",ids.Ids).Error
	return err
}

// @title    UpdatePicture
// @description   update a Picture
// @param     picture          *model.Picture
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdatePicture(picture *model.Picture) (err error) {
	err = global.GVA_DB.Save(picture).Error
	return err
}

// @title    GetPicture
// @description   get the info of a Picture
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Picture        Picture

func GetPicture(id uint) (err error, picture model.Picture) {
	err = global.GVA_DB.Where("id = ?", id).First(&picture).Error
	return
}

// @title    GetPictureInfoList
// @description   get Picture list by pagination, 分页获取Picture
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetPictureInfoList(info request.PictureSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Picture{})
    var pictures []model.Picture
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pictures).Error
	return err, pictures, total
}
