package request

import "gin-vue-admin/model"

type HardSearch struct{
    model.Hard
    PageInfo
}