package autocode

import (
	"raptor/server/global"
	"raptor/server/model/autocode"
	autoCodeReq "raptor/server/model/autocode/request"
	"raptor/server/model/common/request"
)

type ServiceService struct {
}

// CreateService 创建Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService) CreateService(service autocode.Service) (err error) {
	err = global.GVA_DB.Create(&service).Error
	return err
}

// DeleteService 删除Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService)DeleteService(service autocode.Service) (err error) {
	_ = global.GVA_DB.Model(&service).Association("Products").Clear()
	err = global.GVA_DB.Delete(&service).Error
	return err
}

// DeleteServiceByIds 批量删除Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService)DeleteServiceByIds(ids request.IdsReq) (err error) {
	_ = global.GVA_DB.Delete(&[]autocode.Service{},"id in ?",ids.Ids).Association("Products").Clear()
	err = global.GVA_DB.Delete(&[]autocode.Service{},"id in ?",ids.Ids).Error
	return err
}

// UpdateService 更新Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService)UpdateService(service autocode.Service) (err error) {
	err = global.GVA_DB.Save(&service).Error
	return err
}

// GetService 根据id获取Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService)GetService(id uint) (err error, service autocode.Service) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Products").First(&service).Error
	return
}

// GetServiceInfoList 分页获取Service记录
// Author [piexlmax](https://github.com/piexlmax)
func (serviceService *ServiceService)GetServiceInfoList(info autoCodeReq.ServiceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Service{})
    var services []autocode.Service
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Type != "" {
        db = db.Where("type LIKE ?","%"+ info.Type+"%")
    }
    if info.Url != "" {
        db = db.Where("url LIKE ?","%"+ info.Url+"%")
    }
    if info.Branch != "" {
        db = db.Where("branch LIKE ?","%"+ info.Branch+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Preload("Products").Find(&services).Error
	return err, services, total
}
