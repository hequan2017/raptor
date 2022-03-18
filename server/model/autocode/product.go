// 自动生成模板Product
package autocode

import (
	"raptor/server/global"
)

// Product 结构体
// 如果含有time.Time 请自行import time包
type Product struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`
      Other  string `json:"other" form:"other" gorm:"column:other;comment:其他;size:255;"`
}


// TableName Product 表名
func (Product) TableName() string {
  return "product"
}

