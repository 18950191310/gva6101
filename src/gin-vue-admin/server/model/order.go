// 自动生成模板Order
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Order struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(255);size:255;"`
      Mount  int `json:"mount" form:"mount" gorm:"column:mount;comment:数量;type:int;size:10;"`
      Total  float64 `json:"total" form:"total" gorm:"column:total;comment:总价;type:float;"`
      User  string `json:"user" form:"user" gorm:"column:user;comment:用户名;type:varchar(255);size:255;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机;type:varchar(255);size:255;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:收货地址;type:varchar(255);size:255;"`
}


func (Order) TableName() string {
  return "order"
}

