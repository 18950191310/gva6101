// 自动生成模板Register
package model

import (
	"gin-vue-admin/global"
      "github.com/satori/go.uuid"

)

// 如果含有time.Time 请自行import time包
type Register struct {
      global.GVA_MODEL
      UUID      uuid.UUID    `json:"uuid" gorm:"comment:UUID"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:用户名;type:varchar;"`
      NickName  string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;type:varchar;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar;"`
      AuthorityId  string `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:权限;type:varchar;"`
}


func (Register) TableName() string {
  return "register"
}

