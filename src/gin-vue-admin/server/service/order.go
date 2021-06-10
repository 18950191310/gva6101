package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateOrder
//@description: 创建Order记录
//@param: order model.Order
//@return: err error

func CreateOrder(order model.Order) (err error) {
	err = global.GVA_DB.Create(&order).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOrder
//@description: 删除Order记录
//@param: order model.Order
//@return: err error

func DeleteOrder(order model.Order) (err error) {
	err = global.GVA_DB.Delete(&order).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteOrderByIds
//@description: 批量删除Order记录
//@param: ids request.IdsReq
//@return: err error

func DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Order{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateOrder
//@description: 更新Order记录
//@param: order *model.Order
//@return: err error

func UpdateOrder(order model.Order) (err error) {
	err = global.GVA_DB.Save(&order).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOrder
//@description: 根据id获取Order记录
//@param: id uint
//@return: err error, order model.Order

func GetOrder(id uint) (err error, order model.Order) {
	err = global.GVA_DB.Where("id = ?", id).First(&order).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetOrderInfoList
//@description: 分页获取Order记录
//@param: info request.OrderSearch
//@return: err error, list interface{}, total int64

func GetOrderInfoList(info request.OrderSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Order{})
    var orders []model.Order
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	return err, orders, total
}