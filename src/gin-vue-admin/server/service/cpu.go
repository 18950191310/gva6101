package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateCpu
//@description: 创建Cpu记录
//@param: cpu model.Cpu
//@return: err error

func CreateCpu(cpu model.Cpu) (err error) {
	err = global.GVA_DB.Create(&cpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCpu
//@description: 删除Cpu记录
//@param: cpu model.Cpu
//@return: err error

func DeleteCpu(cpu model.Cpu) (err error) {
	err = global.GVA_DB.Delete(&cpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCpuByIds
//@description: 批量删除Cpu记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCpuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Cpu{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCpu
//@description: 更新Cpu记录
//@param: cpu *model.Cpu
//@return: err error

func UpdateCpu(cpu model.Cpu) (err error) {
	err = global.GVA_DB.Save(&cpu).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCpu
//@description: 根据id获取Cpu记录
//@param: id uint
//@return: err error, cpu model.Cpu

func GetCpu(id uint) (err error, cpu model.Cpu) {
	err = global.GVA_DB.Where("id = ?", id).First(&cpu).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCpuInfoList
//@description: 分页获取Cpu记录
//@param: info request.CpuSearch
//@return: err error, list interface{}, total int64

func GetCpuInfoList(info request.CpuSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Cpu{})
    var cpus []model.Cpu
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cpus).Error
	return err, cpus, total
}