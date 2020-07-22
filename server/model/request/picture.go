package request

import "gin-vue-admin/model"

type PictureSearch struct{
    model.Picture
    PageInfo
}