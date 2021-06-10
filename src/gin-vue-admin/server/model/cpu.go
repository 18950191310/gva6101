// 自动生成模板Cpu
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Cpu struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(255);size:255;"`
      Price  float64 `json:"price" form:"price" gorm:"column:price;comment:价格;type:float;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:出售状态;type:tinyint;"`
      Stock  int `json:"stock" form:"stock" gorm:"column:stock;comment:库存;type:int;size:10;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;type:varchar(255);size:255;"`
}


func (Cpu) TableName() string {
  return "cpu"
}

