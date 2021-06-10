package request

import "gin-vue-admin/model"

type CpuSearch struct{
    model.Cpu
    PageInfo
}