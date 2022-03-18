package autocode

import (
	"raptor/server/global"
	"raptor/server/model/autocode"
	autoCodeReq "raptor/server/model/autocode/request"
	"raptor/server/model/common/request"
)

type AssetService struct {
}

// CreateAsset 创建Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService) CreateAsset(asset autocode.Asset) (err error) {
	err = global.GVA_DB.Create(&asset).Error
	return err
}

// DeleteAsset 删除Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService)DeleteAsset(asset autocode.Asset) (err error) {
	err = global.GVA_DB.Delete(&asset).Error
	return err
}

// DeleteAssetByIds 批量删除Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService)DeleteAssetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Asset{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAsset 更新Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService)UpdateAsset(asset autocode.Asset) (err error) {
	err = global.GVA_DB.Save(&asset).Error
	return err
}

// GetAsset 根据id获取Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService)GetAsset(id uint) (err error, asset autocode.Asset) {
	err = global.GVA_DB.Where("id = ?", id).First(&asset).Error
	return
}

// GetAssetInfoList 分页获取Asset记录
// Author [piexlmax](https://github.com/piexlmax)
func (assetService *AssetService)GetAssetInfoList(info autoCodeReq.AssetSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Asset{})
    var assets []autocode.Asset
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Region != "" {
        db = db.Where("region LIKE ?","%"+ info.Region+"%")
    }
    if info.Type != "" {
        db = db.Where("type LIKE ?","%"+ info.Type+"%")
    }
    if info.Brand != "" {
        db = db.Where("brand LIKE ?","%"+ info.Brand+"%")
    }
    if info.Model != "" {
        db = db.Where("model LIKE ?","%"+ info.Model+"%")
    }
    if info.System != "" {
        db = db.Where("system LIKE ?","%"+ info.System+"%")
    }
    if info.Hostname != "" {
        db = db.Where("hostname LIKE ?","%"+ info.Hostname+"%")
    }
    if info.Sn != "" {
        db = db.Where("sn LIKE ?","%"+ info.Sn+"%")
    }
    if info.First != "" {
        db = db.Where("first LIKE ?","%"+ info.First+"%")
    }
    if info.Intranet != "" {
        db = db.Where("intranet LIKE ?","%"+ info.Intranet+"%")
    }
    if info.Public != "" {
        db = db.Where("public LIKE ?","%"+ info.Public+"%")
    }
    if info.Status != "" {
        db = db.Where("status LIKE ?","%"+ info.Status+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&assets).Error
	return err, assets, total
}
