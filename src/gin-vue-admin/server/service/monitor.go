package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMonitor
//@description: 创建Monitor记录
//@param: monitor model.Monitor
//@return: err error

func CreateMonitor(monitor model.Monitor) (err error) {
	err = global.GVA_DB.Create(&monitor).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMonitor
//@description: 删除Monitor记录
//@param: monitor model.Monitor
//@return: err error

func DeleteMonitor(monitor model.Monitor) (err error) {
	err = global.GVA_DB.Delete(&monitor).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMonitorByIds
//@description: 批量删除Monitor记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMonitorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Monitor{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMonitor
//@description: 更新Monitor记录
//@param: monitor *model.Monitor
//@return: err error

func UpdateMonitor(monitor model.Monitor) (err error) {
	err = global.GVA_DB.Save(&monitor).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMonitor
//@description: 根据id获取Monitor记录
//@param: id uint
//@return: err error, monitor model.Monitor

func GetMonitor(id uint) (err error, monitor model.Monitor) {
	err = global.GVA_DB.Where("id = ?", id).First(&monitor).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMonitorInfoList
//@description: 分页获取Monitor记录
//@param: info request.MonitorSearch
//@return: err error, list interface{}, total int64

func GetMonitorInfoList(info request.MonitorSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Monitor{})
    var monitors []model.Monitor
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&monitors).Error
	return err, monitors, total
}