// 自动生成模板Picture
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Picture struct {
      gorm.Model
      ImgPath  string `json:"img_path" form:"img_path" gorm:"column:img_path;comment:'图片路径';type:varchar(255)"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:'图片描述';type:text"`
      CatId  int `json:"cat_id" form:"cat_id" gorm:"column:cat_id;comment:'所属类别id';type:int(11)"`
      DisplayOrder  int `json:"display_order" form:"display_order" gorm:"column:display_order;comment:'排序';type:smallint"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:'状态';type:smallint"`
      UserId  int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'用户id';type:int"` 
}

