// 自动生成模板Recolist
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Recolist struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar;"`
      Cpu  string `json:"cpu" form:"cpu" gorm:"column:cpu;comment:cpu;type:varchar;"`
      Gpu  string `json:"gpu" form:"gpu" gorm:"column:gpu;comment:显卡;type:varchar;"`
      Storage  string `json:"storage" form:"storage" gorm:"column:storage;comment:内存;type:varchar;"`
      Hard  string `json:"hard" form:"hard" gorm:"column:hard;comment:硬盘;type:varchar;"`
      Monitor  string `json:"monitor" form:"monitor" gorm:"column:monitor;comment:显示器;type:varchar;"`
      Total  float64 `json:"total" form:"total" gorm:"column:total;comment:总价;type:float;"`
}


func (Recolist) TableName() string {
  return "recolist"
}

