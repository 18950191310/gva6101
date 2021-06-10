// 自动生成模板CuserDetail
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type CuserDetail struct {
      global.GVA_MODEL
      Username    string       `json:"userName" gorm:"comment:用户登录名"`
      NickName  string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;type:varchar;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;type:varchar;"`
      Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:性别;type:varchar;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:称呼;type:varchar;"`
      Region  string `json:"region" form:"region" gorm:"column:region;comment:地区;type:varchar;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:地址;type:varchar;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar;"`
      Wechat  string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:微信号;type:varchar;"`
      Number  string `json:"number" form:"number" gorm:"column:number;comment:编号;type:varchar;"`
}


func (CuserDetail) TableName() string {
  return "cuserDetail"
}

