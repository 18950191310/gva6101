package request

import "gin-vue-admin/model"

type RecolistSearch struct{
    model.Recolist
    PageInfo
}