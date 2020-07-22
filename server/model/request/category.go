package request

import "gin-vue-admin/model"

type CategorySearch struct{
    model.Category
    PageInfo
}