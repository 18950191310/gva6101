package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStorage
//@description: 创建Storage记录
//@param: storage model.Storage
//@return: err error

func CreateStorage(storage model.Storage) (err error) {
	err = global.GVA_DB.Create(&storage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStorage
//@description: 删除Storage记录
//@param: storage model.Storage
//@return: err error

func DeleteStorage(storage model.Storage) (err error) {
	err = global.GVA_DB.Delete(&storage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStorageByIds
//@description: 批量删除Storage记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStorageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Storage{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStorage
//@description: 更新Storage记录
//@param: storage *model.Storage
//@return: err error

func UpdateStorage(storage model.Storage) (err error) {
	err = global.GVA_DB.Save(&storage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStorage
//@description: 根据id获取Storage记录
//@param: id uint
//@return: err error, storage model.Storage

func GetStorage(id uint) (err error, storage model.Storage) {
	err = global.GVA_DB.Where("id = ?", id).First(&storage).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStorageInfoList
//@description: 分页获取Storage记录
//@param: info request.StorageSearch
//@return: err error, list interface{}, total int64

func GetStorageInfoList(info request.StorageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Storage{})
    var storages []model.Storage
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&storages).Error
	return err, storages, total
}