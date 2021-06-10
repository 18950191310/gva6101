package request

import "gin-vue-admin/model"

type CuserDetailSearch struct{
    model.CuserDetail
    PageInfo
}