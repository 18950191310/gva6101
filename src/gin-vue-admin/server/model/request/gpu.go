package request

import "gin-vue-admin/model"

type GpuSearch struct{
    model.Gpu
    PageInfo
}