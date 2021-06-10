package request

import "gin-vue-admin/model"

type MonitorSearch struct{
    model.Monitor
    PageInfo
}