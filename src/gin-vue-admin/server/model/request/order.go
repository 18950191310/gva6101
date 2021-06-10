package request

import "gin-vue-admin/model"

type OrderSearch struct{
    model.Order
    PageInfo
}