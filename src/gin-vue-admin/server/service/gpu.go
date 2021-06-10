package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateGpu
//@description: 创建Gpu记录
//@param: gpu model.Gpu
//@return: err error

func CreateGpu(gpu model.Gpu) (err error) {
	err = global.GVA_DB.Create(&gpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGpu
//@description: 删除Gpu记录
//@param: gpu model.Gpu
//@return: err error

func DeleteGpu(gpu model.Gpu) (err error) {
	err = global.GVA_DB.Delete(&gpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGpuByIds
//@description: 批量删除Gpu记录
//@param: ids request.IdsReq
//@return: err error

func DeleteGpuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Gpu{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateGpu
//@description: 更新Gpu记录
//@param: gpu *model.Gpu
//@return: err error

func UpdateGpu(gpu model.Gpu) (err error) {
	err = global.GVA_DB.Save(&gpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGpu
//@description: 根据id获取Gpu记录
//@param: id uint
//@return: err error, gpu model.Gpu

func GetGpu(id uint) (err error, gpu model.Gpu) {
	err = global.GVA_DB.Where("id = ?", id).First(&gpu).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGpuInfoList
//@description: 分页获取Gpu记录
//@param: info request.GpuSearch
//@return: err error, list interface{}, total int64

func GetGpuInfoList(info request.GpuSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Gpu{})
    var gpus []model.Gpu
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&gpus).Error
	return err, gpus, total
}