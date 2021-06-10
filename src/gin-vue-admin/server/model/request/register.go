package request

import "gin-vue-admin/model"

type RegisterSearch struct{
    model.Register
    PageInfo
}