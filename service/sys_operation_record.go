package service

import (
	"gin-learn/global"
	"gin-learn/model"
	"gin-learn/model/request"
)

// 创建操作记录
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.G_DB.Create(&sysOperationRecord).Error
	return err
}

// 批量删除记录
func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.G_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// 删除操作记录
func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.G_DB.Delete(&sysOperationRecord).Error
	return err
}

// 根据id获取单条操作记录
func GetSysOperationRecord(id uint) (err error, sysOperationRecord model.SysOperationRecord) {
	err = global.G_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

// 分页获取操作记录列表
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.G_DB.Model(&model.SysOperationRecord{})
	var sysOperationRecords []model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
