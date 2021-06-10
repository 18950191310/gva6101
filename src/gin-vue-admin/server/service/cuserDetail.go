package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateCuserDetail
//@description: 创建CuserDetail记录
//@param: cuserDetail model.CuserDetail
//@return: err error

func CreateCuserDetail(cuserDetail model.CuserDetail) (err error) {
	err = global.GVA_DB.Create(&cuserDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCuserDetail
//@description: 删除CuserDetail记录
//@param: cuserDetail model.CuserDetail
//@return: err error

func DeleteCuserDetail(cuserDetail model.CuserDetail) (err error) {
	err = global.GVA_DB.Delete(&cuserDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCuserDetailByIds
//@description: 批量删除CuserDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCuserDetailByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CuserDetail{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCuserDetail
//@description: 更新CuserDetail记录
//@param: cuserDetail *model.CuserDetail
//@return: err error

func UpdateCuserDetail(cuserDetail model.CuserDetail) (err error) {
	err = global.GVA_DB.Save(&cuserDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCuserDetail
//@description: 根据id获取CuserDetail记录
//@param: id uint
//@return: err error, cuserDetail model.CuserDetail

func GetCuserDetail(id uint) (err error, cuserDetail model.CuserDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&cuserDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCuserDetailInfoList
//@description: 分页获取CuserDetail记录
//@param: info request.CuserDetailSearch
//@return: err error, list interface{}, total int64

func GetCuserDetailInfoList(info request.CuserDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.CuserDetail{})
    var cuserDetails []model.CuserDetail
	if info.Username != "" {
		db = db.Where("`username` = ?",info.Username)
	}
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.NickName != "" {
        db = db.Where("`nick_name` LIKE ?","%"+ info.NickName+"%")
    }
    if info.Phone != "" {
        db = db.Where("`phone` LIKE ?","%"+ info.Phone+"%")
    }
    if info.Sex != "" {
        db = db.Where("`sex` = ?",info.Sex)
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Region != "" {
        db = db.Where("`region` LIKE ?","%"+ info.Region+"%")
    }
    if info.Address != "" {
        db = db.Where("`address` LIKE ?","%"+ info.Address+"%")
    }
    if info.Password != "" {
        db = db.Where("`password` = ?",info.Password)
    }
    if info.Wechat != "" {
        db = db.Where("`wechat` LIKE ?","%"+ info.Wechat+"%")
    }
    if info.Number != "" {
        db = db.Where("`number` LIKE ?","%"+ info.Number+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cuserDetails).Error
	return err, cuserDetails, total
}