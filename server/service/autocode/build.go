package autocode

import (
	"raptor/server/global"
	"raptor/server/model/autocode"
	autoCodeReq "raptor/server/model/autocode/request"
	"raptor/server/model/common/request"
)

type BuildService struct {
}

// CreateBuild 创建Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService) CreateBuild(build autocode.Build) (err error) {
	err = global.GVA_DB.Create(&build).Error
	return err
}

// DeleteBuild 删除Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService)DeleteBuild(build autocode.Build) (err error) {
	err = global.GVA_DB.Delete(&build).Error
	return err
}

// DeleteBuildByIds 批量删除Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService)DeleteBuildByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Build{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBuild 更新Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService)UpdateBuild(build autocode.Build) (err error) {
	err = global.GVA_DB.Save(&build).Error
	return err
}

// GetBuild 根据id获取Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService)GetBuild(id uint) (err error, build autocode.Build) {
	err = global.GVA_DB.Where("id = ?", id).First(&build).Error
	return
}

// GetBuildInfoList 分页获取Build记录
// Author [piexlmax](https://github.com/piexlmax)
func (buildService *BuildService)GetBuildInfoList(info autoCodeReq.BuildSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Build{})
    var builds []autocode.Build
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Static != "" {
        db = db.Where("static LIKE ?","%"+ info.Static+"%")
    }
    if info.Version != "" {
        db = db.Where("version LIKE ?","%"+ info.Version+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&builds).Error
	return err, builds, total
}
