// 自动生成模板Service
package autocode

import (
	"raptor/server/global"
)

// Service 结构体
// 如果含有time.Time 请自行import time包
type Service struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:项目名称;size:255;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:库地址;size:255;"`
      Branch  string `json:"branch" form:"branch" gorm:"column:branch;comment:分支;size:255;"`
      Start  string `json:"start" form:"start" gorm:"column:start;comment:启动命令;size:255;"`
      Stop  string `json:"stop" form:"stop" gorm:"column:stop;comment:关闭命令;size:255;"`
      Activity  string `json:"activity" form:"activity" gorm:"column:activity;comment:探活地址;size:255;"`
      Env  string `json:"env" form:"env" gorm:"column:env;comment:环境变量;"`
      Shell  string `json:"shell" form:"shell" gorm:"column:shell;comment:打包命令;"`
      Other  string `json:"other" form:"other" gorm:"column:other;comment:其他;"`
      Products  []Product  `json:"products" gorm:"many2many:service_product_id;"`
}


// TableName Service 表名
func (Service) TableName() string {
  return "service"
}

