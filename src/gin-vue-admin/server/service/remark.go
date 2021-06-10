package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRemark
//@description: 创建Remark记录
//@param: remark model.Remark
//@return: err error

func CreateRemark(remark model.Remark) (err error) {
	err = global.GVA_DB.Create(&remark).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRemark
//@description: 删除Remark记录
//@param: remark model.Remark
//@return: err error

func DeleteRemark(remark model.Remark) (err error) {
	err = global.GVA_DB.Delete(&remark).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRemarkByIds
//@description: 批量删除Remark记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRemarkByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Remark{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRemark
//@description: 更新Remark记录
//@param: remark *model.Remark
//@return: err error

func UpdateRemark(remark model.Remark) (err error) {
	err = global.GVA_DB.Save(&remark).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRemark
//@description: 根据id获取Remark记录
//@param: id uint
//@return: err error, remark model.Remark

func GetRemark(id uint) (err error, remark model.Remark) {
	err = global.GVA_DB.Where("id = ?", id).First(&remark).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRemarkInfoList
//@description: 分页获取Remark记录
//@param: info request.RemarkSearch
//@return: err error, list interface{}, total int64

func GetRemarkInfoList(info request.RemarkSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Remark{})
    var remarks []model.Remark
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
        db = db.Where("`username` LIKE ?","%"+ info.Username+"%")
    }
    if info.Comment != "" {
        db = db.Where("`comment` LIKE ?","%"+ info.Comment+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&remarks).Error
	return err, remarks, total
}