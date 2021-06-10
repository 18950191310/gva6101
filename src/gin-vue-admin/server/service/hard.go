package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateHard
//@description: 创建Hard记录
//@param: hard model.Hard
//@return: err error

func CreateHard(hard model.Hard) (err error) {
	err = global.GVA_DB.Create(&hard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHard
//@description: 删除Hard记录
//@param: hard model.Hard
//@return: err error

func DeleteHard(hard model.Hard) (err error) {
	err = global.GVA_DB.Delete(&hard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHardByIds
//@description: 批量删除Hard记录
//@param: ids request.IdsReq
//@return: err error

func DeleteHardByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Hard{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateHard
//@description: 更新Hard记录
//@param: hard *model.Hard
//@return: err error

func UpdateHard(hard model.Hard) (err error) {
	err = global.GVA_DB.Save(&hard).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHard
//@description: 根据id获取Hard记录
//@param: id uint
//@return: err error, hard model.Hard

func GetHard(id uint) (err error, hard model.Hard) {
	err = global.GVA_DB.Where("id = ?", id).First(&hard).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHardInfoList
//@description: 分页获取Hard记录
//@param: info request.HardSearch
//@return: err error, list interface{}, total int64

func GetHardInfoList(info request.HardSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Hard{})
    var hards []model.Hard
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hards).Error
	return err, hards, total
}