package autocode

import (
	"raptor/server/global"
	"raptor/server/model/autocode"
	autoCodeReq "raptor/server/model/autocode/request"
	"raptor/server/model/common/request"
)

type KeyService struct {
}

// CreateKey 创建Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService) CreateKey(key autocode.Key) (err error) {
	err = global.GVA_DB.Create(&key).Error
	return err
}

// DeleteKey 删除Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService)DeleteKey(key autocode.Key) (err error) {
	err = global.GVA_DB.Delete(&key).Error
	return err
}

// DeleteKeyByIds 批量删除Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService)DeleteKeyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Key{},"id in ?",ids.Ids).Error
	return err
}

// UpdateKey 更新Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService)UpdateKey(key autocode.Key) (err error) {
	err = global.GVA_DB.Save(&key).Error
	return err
}

// GetKey 根据id获取Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService)GetKey(id uint) (err error, key autocode.Key) {
	err = global.GVA_DB.Where("id = ?", id).First(&key).Error
	return
}

// GetKeyInfoList 分页获取Key记录
// Author [piexlmax](https://github.com/piexlmax)
func (keyService *KeyService)GetKeyInfoList(info autoCodeReq.KeySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Key{})
    var keys []autocode.Key
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Type != "" {
        db = db.Where("type LIKE ?","%"+ info.Type+"%")
    }
    if info.Region != "" {
        db = db.Where("region LIKE ?","%"+ info.Region+"%")
    }
    if info.Keyid != "" {
        db = db.Where("keyid LIKE ?","%"+ info.Keyid+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&keys).Error
	return err, keys, total
}
