// 自动生成模板Category
package model

import (
	"github.com/jinzhu/gorm"
)

// 如果含有time.Time 请自行import time包
type Category struct {
      gorm.Model
      CatName  string `json:"cat_name" form:"cat_name" gorm:"column:cat_name;comment:'类别名称';type:varchar(255)"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:'类别描述';type:text"`
      Pid  int `json:"pid" form:"pid" gorm:"column:pid;comment:'父id';type:int(11)"`
      DisplayOrder  int `json:"display_order" form:"display_order" gorm:"column:display_order;comment:'排序';type:smallint(4)"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:'状态';type:smallint(4)"`
      UserId  int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:'用户id';type:int(11)"` 
}

