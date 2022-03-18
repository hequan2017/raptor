// 自动生成模板Asset
package autocode

import (
      "gorm.io/gorm"
      "time"
)

// Asset 结构体
// 如果含有time.Time 请自行import time包
type Asset struct {
      ID        uint           `gorm:"primarykey"` // 主键ID
      CreatedAt time.Time      // 创建时间
      UpdatedAt time.Time      // 更新时间
      DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
      Region    string     `json:"region" form:"region" gorm:"column:region;comment:区域;size:255;"`
      Type      string     `json:"type" form:"type" gorm:"column:type;comment:类型;size:255;"`
      Brand     string     `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;size:255;"`
      Model     string     `json:"model" form:"model" gorm:"column:model;comment:型号;size:255;"`
      System    string     `json:"system" form:"system" gorm:"column:system;comment:系统;size:255;"`
      Hostname  string     `json:"hostname" form:"hostname" gorm:"column:hostname;comment:主机名;size:255;"`
      Sn        string     `json:"sn" form:"sn" gorm:"column:sn;comment:资产编号;size:255;"`
      First     string     `json:"first" form:"first" gorm:"column:first;comment:一级负责人;size:255;"`
      Intranet  string     `json:"intranet" form:"intranet" gorm:"column:intranet;comment:内网IP;size:255;"`
      Public    string     `json:"public" form:"public" gorm:"column:public;comment:外网IP;size:255;"`
      Status    string     `json:"status" form:"status" gorm:"column:status;comment:状态;size:255;"`
      Cpu       string     `json:"cpu" form:"cpu" gorm:"column:cpu;comment:CPU;size:255;"`
      Memory    string     `json:"memory" form:"memory" gorm:"column:memory;comment:内存;size:255;"`
      Disk      string     `json:"disk" form:"disk" gorm:"column:disk;comment:硬盘;size:255;"`
      Other     string     `json:"other" form:"other" gorm:"column:other;comment:其他;size:255;"`
      Uptime    *time.Time `json:"uptime" form:"uptime" gorm:"column:uptime;comment:上线时间;"`
      Products  []Product  `json:"products" gorm:"many2many:asset_product_id;"`
}


// TableName Asset 表名
func (Asset) TableName() string {
  return "asset"
}

