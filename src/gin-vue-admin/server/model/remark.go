// 自动生成模板Remark
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Remark struct {
      global.GVA_MODEL
      Username  string `json:"username" form:"username" gorm:"column:username;comment:用户名;type:varchar(255);size:255;"`
      Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;type:varchar(255);size:255;"`
}


func (Remark) TableName() string {
  return "remark"
}

