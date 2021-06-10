package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRecolist
//@description: 创建Recolist记录
//@param: recolist model.Recolist
//@return: err error

func CreateRecolist(recolist model.Recolist) (err error) {
	err = global.GVA_DB.Create(&recolist).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRecolist
//@description: 删除Recolist记录
//@param: recolist model.Recolist
//@return: err error

func DeleteRecolist(recolist model.Recolist) (err error) {
	err = global.GVA_DB.Delete(&recolist).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRecolistByIds
//@description: 批量删除Recolist记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRecolistByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Recolist{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRecolist
//@description: 更新Recolist记录
//@param: recolist *model.Recolist
//@return: err error

func UpdateRecolist(recolist model.Recolist) (err error) {
	err = global.GVA_DB.Save(&recolist).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRecolist
//@description: 根据id获取Recolist记录
//@param: id uint
//@return: err error, recolist model.Recolist

func GetRecolist(id uint) (err error, recolist model.Recolist) {
	err = global.GVA_DB.Where("id = ?", id).First(&recolist).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRecolistInfoList
//@description: 分页获取Recolist记录
//@param: info request.RecolistSearch
//@return: err error, list interface{}, total int64

func GetRecolistInfoList(info request.RecolistSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Recolist{})
    var recolists []model.Recolist
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Cpu != "" {
        db = db.Where("`cpu` LIKE ?","%"+ info.Cpu+"%")
    }
    if info.Gpu != "" {
        db = db.Where("`gpu` LIKE ?","%"+ info.Gpu+"%")
    }
    if info.Storage != "" {
        db = db.Where("`storage` LIKE ?","%"+ info.Storage+"%")
    }
    if info.Hard != "" {
        db = db.Where("`hard` LIKE ?","%"+ info.Hard+"%")
    }
    if info.Monitor != "" {
        db = db.Where("`monitor` LIKE ?","%"+ info.Monitor+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&recolists).Error
	return err, recolists, total
}