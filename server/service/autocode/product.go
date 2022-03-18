package autocode

import (
	"raptor/server/global"
	"raptor/server/model/autocode"
	"raptor/server/model/common/request"
    autoCodeReq "raptor/server/model/autocode/request"
)

type ProductService struct {
}

// CreateProduct 创建Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) CreateProduct(product autocode.Product) (err error) {
	err = global.GVA_DB.Create(&product).Error
	return err
}

// DeleteProduct 删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService)DeleteProduct(product autocode.Product) (err error) {
	err = global.GVA_DB.Delete(&product).Error
	return err
}

// DeleteProductByIds 批量删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService)DeleteProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Product{},"id in ?",ids.Ids).Error
	return err
}

// UpdateProduct 更新Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService)UpdateProduct(product autocode.Product) (err error) {
	err = global.GVA_DB.Save(&product).Error
	return err
}

// GetProduct 根据id获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService)GetProduct(id uint) (err error, product autocode.Product) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return
}

// GetProductInfoList 分页获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService)GetProductInfoList(info autoCodeReq.ProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Product{})
    var products []autocode.Product
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&products).Error
	return err, products, total
}
