package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRegister
//@description: 创建Register记录
//@param: register model.Register
//@return: err error

func CreateRegister(register model.Register) (err error) {
	err = global.GVA_DB.Create(&register).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRegister
//@description: 删除Register记录
//@param: register model.Register
//@return: err error

func DeleteRegister(register model.Register) (err error) {
	err = global.GVA_DB.Delete(&register).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRegisterByIds
//@description: 批量删除Register记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRegisterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Register{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRegister
//@description: 更新Register记录
//@param: register *model.Register
//@return: err error

func UpdateRegister(register model.Register) (err error) {
	err = global.GVA_DB.Save(&register).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRegister
//@description: 根据id获取Register记录
//@param: id uint
//@return: err error, register model.Register

func GetRegister(id uint) (err error, register model.Register) {
	err = global.GVA_DB.Where("id = ?", id).First(&register).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRegisterInfoList
//@description: 分页获取Register记录
//@param: info request.RegisterSearch
//@return: err error, list interface{}, total int64

func GetRegisterInfoList(info request.RegisterSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Register{})
    var registers []model.Register
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
        db = db.Where("`username` LIKE ?","%"+ info.Username+"%")
    }
    if info.NickName != "" {
        db = db.Where("`nick_name` LIKE ?","%"+ info.NickName+"%")
    }
    if info.Password != "" {
        db = db.Where("`password` LIKE ?","%"+ info.Password+"%")
    }
    if info.AuthorityId != "" {
        db = db.Where("`authority_id` LIKE ?","%"+ info.AuthorityId+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&registers).Error
	return err, registers, total
}