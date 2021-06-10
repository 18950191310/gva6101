package request

import "gin-vue-admin/model"

type RemarkSearch struct{
    model.Remark
    PageInfo
}