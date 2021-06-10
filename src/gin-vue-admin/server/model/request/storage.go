package request

import "gin-vue-admin/model"

type StorageSearch struct{
    model.Storage
    PageInfo
}